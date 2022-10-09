package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
    AuthorizationToken string                                `json:"authorizationToken"`
	RequestContext     events.AppSyncLambdaAuthorizerRequestContext `json:"requestContext"`
}

type JWTPayload struct {
	UserId string `json:"userId"`
}

func HandleRequest(ctx context.Context, event Event) (*events.AppSyncLambdaAuthorizerResponse, error) {
		jwt := event.AuthorizationToken
		authenticationEndpoint := fmt.Sprintf("%s/auth/verify", os.Getenv("SERVICE_API_ENDPOINT"))
		body := struct {
			Token string `json:"token"`
		}{
			Token: jwt,
		}
		encodedPayload, err := json.Marshal(&body)
		if err != nil {
			return generateUnauthorizedResponse(), nil
		}
		// // send REST API to verify jwt
		resp, err := http.Post(authenticationEndpoint, "application/json", bytes.NewBuffer(encodedPayload))
		if err != nil {
			return generateUnauthorizedResponse(), nil
		}

		payload := struct {
			UId string `json:"uid"`
			Exp int `json:"exp"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return generateUnauthorizedResponse(), nil
		}
		return generateAuthorizedResponse(payload.UId), nil
}

func generateUnauthorizedResponse() *events.AppSyncLambdaAuthorizerResponse {
	return &events.AppSyncLambdaAuthorizerResponse{
		IsAuthorized: false,
	}
}

func generateAuthorizedResponse(userId string) *events.AppSyncLambdaAuthorizerResponse {
	return &events.AppSyncLambdaAuthorizerResponse{
		IsAuthorized: true,
		ResolverContext: map[string]interface{}{
			userId: userId,
		},
	}
}

func main() {
        lambda.Start(HandleRequest)
}