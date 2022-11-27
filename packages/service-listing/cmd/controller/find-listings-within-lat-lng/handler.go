package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-listing/internal/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	repo := repository.NewListingRepository(db)
	repo.Setup()
	service := service.NewListingService(repo)
	apigateway := controller.NewApiGatewayHandler(service)
	return apigateway.FindListingsWithinLatLng(request)
}

func main() {
	lambda.Start(handler)
}
