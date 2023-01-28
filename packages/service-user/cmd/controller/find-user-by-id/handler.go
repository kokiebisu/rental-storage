package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-user/internal"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-user/internal/helper"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller, err := controller.New()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	user, err := controller.FindUserById(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := helper.Stringify(user)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
