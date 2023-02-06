package responses

import (
	"github.com/aws/aws-lambda-go/events"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
	"github.com/kokiebisu/rental-storage/service-listing/internal/helper"
)

func SendFailureResponse(err *errors.CustomError) (events.APIGatewayProxyResponse, error) {
	payload := err.GetPayload()
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: int(err.StatusCode),
		Body:       result,
	}, nil
}

func SendSuccessResponse(payload interface{}) (events.APIGatewayProxyResponse, error) {
	if payload == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
		}, nil
	}
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}
