package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/sender"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-user/internal/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	repo := repository.NewUserRepository(db)
	eventSender := sender.New()
	repo.Setup()
	service := service.NewUserService(repo, eventSender)
	apigateway := controller.New(service)
	return apigateway.RemoveUserById(request)
}

func main() {
	lambda.Start(handler)
}
