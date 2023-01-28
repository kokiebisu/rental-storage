package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	responses "github.com/kokiebisu/rental-storage/service-booking/internal"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/controller"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	apigateway := controller.New()
	bookingId, err := apigateway.CreateBooking(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(bookingId)
}

func main() {
	lambda.Start(handler)
}
