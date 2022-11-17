package main

import (
	"service-authentication/internal/adapter/controller"
	"service-authentication/internal/core/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	service := service.NewEncryptionService()
	controller := controller.New(service)
	return controller.SignUp(request)
}

func main() {
	lambda.Start(handler)
}
