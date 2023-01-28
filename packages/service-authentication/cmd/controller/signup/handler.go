package main

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"

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
