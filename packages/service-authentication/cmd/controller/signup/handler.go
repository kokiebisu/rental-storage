package main

import (
	responses "github.com/kokiebisu/rental-storage/service-authentication/internal"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller := controller.New()
	payload, err := controller.SignUp(request)
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
