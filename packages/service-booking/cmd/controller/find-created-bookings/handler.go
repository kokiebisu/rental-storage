package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	responses "github.com/kokiebisu/rental-storage/service-booking/internal"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-booking/internal/helper"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	apigateway := controller.New()
	bookings, err := apigateway.FindUserBookings(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := helper.Stringify(bookings)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
