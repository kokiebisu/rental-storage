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

	"service-authorizer/pkg/port"
)

func HandleRequest(ctx context.Context, event port.Event) (*events.AppSyncLambdaAuthorizerResponse, error) {
	jwt := event.AuthorizationToken
	body := struct {
		Token string `json:"AuthorizationToken"`
	}{
		Token: jwt,
	}
	encodedPayload, err := json.Marshal(&body)
	if err != nil {
		return generateUnauthorizedResponse(), nil
	}
	authenticationEndpoint := fmt.Sprintf("%s/auth/verify", os.Getenv("SERVICE_API_ENDPOINT"))
	// send REST API to verify jwt
	resp, err := http.Post(authenticationEndpoint, "application/json", bytes.NewBuffer(encodedPayload))
	if err != nil {
		return generateUnauthorizedResponse(), nil
	}
	payload := struct {
		UId string `json:"uid"`
		Exp int    `json:"exp"`
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

func generateAuthorizedResponse(uid string) *events.AppSyncLambdaAuthorizerResponse {
	return &events.AppSyncLambdaAuthorizerResponse{
		IsAuthorized: true,
		ResolverContext: map[string]interface{}{
			"uid": uid,
		},
	}
}

func main() {
	lambda.Start(HandleRequest)
}
