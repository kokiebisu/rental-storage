package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
)

// checks if the authorizationToken in the payload is valid
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	service := service.NewEncryptionService()
	controller := controller.New(service)
	return controller.Verify(request)
}

func main() {
	lambda.Start(handler)
}
