package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func SendResponse(token string) (events.APIGatewayProxyResponse, error) {
	payload := &struct{
		Token string `json:"token"`
	}{
		Token: token,
	}
	encoded, _ := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		Body: string(encoded), 
		StatusCode: 200,
	}, nil
}