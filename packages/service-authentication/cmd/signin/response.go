package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func SendResponse(token string) (events.APIGatewayProxyResponse, error) {
	payload := &struct{
		AuthorizationToken string `json:"authorizationToken"`
	}{
		AuthorizationToken: token,
	}
	encoded, _ := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		Body: string(encoded), 
		StatusCode: 200,
	}, nil
}