package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-space/cmd/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c, err := controller.New()
	if err != nil {
		return controller.SendFailureResponse(err)
	}
	space, err := c.DeleteSpaceById(request)
	if err != nil {
		return controller.SendFailureResponse(err)
	}
	return controller.SendSuccessResponse(space)
}

func main() {
	lambda.Start(handler)
}
