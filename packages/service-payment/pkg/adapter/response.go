package adapter

import (
	"encoding/json"

	"service-payment/pkg/port"

	"github.com/aws/aws-lambda-go/events"
)

func SendResponse(payload interface{}) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(payload)
	if err != nil {
		error := &port.Error{
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