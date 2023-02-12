package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-space/internal"
	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c, err := controller.New()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	payload, err := c.FindSpaces(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(payload)
}

func main() {
	lambda.Start(handler)
}