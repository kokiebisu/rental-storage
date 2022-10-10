package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type VerifyPayload struct {
	AuthorizationToken      string `json:"token"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bodyRequest := VerifyPayload{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	claims, err := verifyJWT(bodyRequest.AuthorizationToken)
	if err != nil {
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
    }
	
	encoded, err := json.Marshal(claims)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(encoded), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}