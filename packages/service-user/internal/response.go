package responses

import (
	"github.com/aws/aws-lambda-go/events"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/internal/helper"
)

func SendFailureResponse(err *errors.CustomError) (events.APIGatewayProxyResponse, error) {
	payload := err.GetPayload()
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: int(err.StatusCode),
		Body:       result,
	}, nil
}

func SendSuccessResponse(data string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}
