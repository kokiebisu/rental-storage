package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kokiebisu/rental-storage/service-authentication/cmd/controller"
)

// checks if the authorizationToken in the payload is valid
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := controller.New()
	payload, err := c.Verify(request)
	if err != nil {
		return controller.SendFailureResponse(err)
	}
	return controller.SendSuccessResponse(payload)
}

func main() {
	lambda.Start(handler)
}
