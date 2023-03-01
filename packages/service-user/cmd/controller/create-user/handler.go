package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kokiebisu/rental-storage/service-user/cmd/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c, err := controller.New()
	if err != nil {
		return controller.SendFailureResponse(err)
	}
	payload, err := c.CreateUser(request)
	if err != nil {
		return controller.SendFailureResponse(err)
	}
	return controller.SendSuccessResponse(payload)
}

func main() {
	lambda.Start(handler)
}
