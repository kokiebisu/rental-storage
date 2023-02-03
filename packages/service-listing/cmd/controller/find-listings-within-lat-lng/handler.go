package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-listing/internal"
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c, err := controller.New()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := c.FindListingsWithinLatLng(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
