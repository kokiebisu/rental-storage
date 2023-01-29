package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-authentication/internal"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller := controller.New()
	payload, err := controller.SignIn(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := helper.Stringify(payload)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
