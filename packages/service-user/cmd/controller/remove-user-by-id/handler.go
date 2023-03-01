package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-user/internal"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller, err := adapter.NewControllerAdapter()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	payload, err := controller.RemoveUserById(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(payload)
}

func main() {
	lambda.Start(handler)
}
