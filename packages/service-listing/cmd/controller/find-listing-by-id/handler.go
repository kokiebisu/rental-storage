package main

import (
	"log"

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
	err = repo.Setup()
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewListingService(repo)
	apigateway := controller.NewApiGatewayHandler(service)
	return apigateway.FindListingById(request)
}

func main() {
	lambda.Start(handler)
}
