package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func SendResponse(payload interface{}) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(payload)
	if err != nil {
		error := &Error{
			Message: "encoding failed",
		}
		errEncoded, _ := json.Marshal(error)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: string(errEncoded),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(encoded),
	}, nil
}