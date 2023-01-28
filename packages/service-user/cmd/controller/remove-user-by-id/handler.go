package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-user/internal"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller, err := controller.New()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := controller.CreateUser(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
