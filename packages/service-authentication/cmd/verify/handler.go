package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"service-authentication/pkg/adapter"
	"service-authentication/pkg/port"
)

// checks if the authorizationToken in the payload is valid
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bodyRequest := port.AuthorizationTokenPayload{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	service := adapter.SetupEncryptionService()
	claims, err := service.VerifyJWT(bodyRequest.AuthorizationToken)
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