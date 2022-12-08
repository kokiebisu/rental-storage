package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-booking/internal/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	repo := repository.NewBookingRepository(db)
	service := service.NewBookingService(repo)
	apigateway := controller.New(service)
	return apigateway.FindUserBookings(request)
}

func main() {
	lambda.Start(handler)
}
