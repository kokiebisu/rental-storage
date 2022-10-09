package main

import (
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
		authenticationEndpoint := fmt.Sprintf("%s/auth/verify?auth=%s", os.Getenv("SERVICE_ENDPOINT"), jwt)
		fmt.Println("ENDPOINT", authenticationEndpoint)
		// // send REST API to verify jwt
		resp, err := http.Get(authenticationEndpoint)
		if err != nil {
			return generateUnauthorizedResponse(), nil
		}
		fmt.Println("RESP", resp)
		payload := JWTPayload{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return generateUnauthorizedResponse(), nil
		}
		fmt.Println("PAYLOAD", payload)

		return generateAuthorizedResponse(payload.UserId), nil
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