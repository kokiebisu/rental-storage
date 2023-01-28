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

	responses "github.com/kokiebisu/rental-storage/service-authorizer/internal"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/port"
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
		return responses.SendUnAuthorizedResponse(), nil
	}
	authenticationEndpoint := fmt.Sprintf("%s/auth/verify", os.Getenv("SERVICE_API_ENDPOINT"))
	// send REST API to verify jwt
	resp, err := http.Post(authenticationEndpoint, "application/json", bytes.NewBuffer(encodedPayload))
	if err != nil {
		return responses.SendUnAuthorizedResponse(), nil
	}
	payload := struct {
		UId string `json:"uid"`
		Exp int    `json:"exp"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return responses.SendUnAuthorizedResponse(), nil
	}
	return responses.SendAuthorizedResponse(payload.UId), nil
}

func main() {
	lambda.Start(HandleRequest)
}
