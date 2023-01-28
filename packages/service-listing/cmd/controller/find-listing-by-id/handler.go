package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	responses "github.com/kokiebisu/rental-storage/service-listing/internal"
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-listing/internal/helper"
	"github.com/kokiebisu/rental-storage/service-listing/internal/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := db.New()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	repo := repository.NewListingRepository(db)
	err = repo.Setup()
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	service := service.NewListingService(repo)
	apigateway := controller.NewApiGatewayHandler(service)
	listing, err := apigateway.FindListingById(request)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	result, err := helper.Stringify(listing)
	if err != nil {
		return responses.SendFailureResponse(err)
	}
	return responses.SendSuccessResponse(result)
}

func main() {
	lambda.Start(handler)
}
