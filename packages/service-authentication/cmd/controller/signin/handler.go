package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"service-authentication/internal/adapter/controller"
	"service-authentication/internal/core/service"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	service := service.NewEncryptionService()
	controller := controller.New(service)
	return controller.SignIn(request)
}

func main() {
	lambda.Start(handler)
}
