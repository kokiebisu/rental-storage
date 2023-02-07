package responses

import (
	"github.com/aws/aws-lambda-go/events"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/internal/helper"
)

func SendFailureResponse(err *customerror.CustomError) (events.APIGatewayProxyResponse, error) {
	payload := err.GetPayload()
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: int(err.StatusCode),
		Body:       result,
	}, nil
}

func SendSuccessResponse(payload interface{}) (events.APIGatewayProxyResponse, error) {
	result, _ := helper.Stringify(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}
