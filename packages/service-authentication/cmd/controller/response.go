package controller

import (
	"github.com/aws/aws-lambda-go/events"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

func SendFailureResponse(err *customerror.CustomError) (events.APIGatewayProxyResponse, error) {
	payload := err.GetPayload()
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: int(err.StatusCode),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: result,
	}, nil
}

func SendSuccessResponse(payload interface{}) (events.APIGatewayProxyResponse, error) {
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: result,
	}, nil
}
