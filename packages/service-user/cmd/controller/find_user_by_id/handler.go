package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-user/internal/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := adapter.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewUserRepository(db)
	repo.Setup()
	service := service.NewUserService(repo)
	apigateway := adapter.NewApiGatewayHandler(service)
	return apigateway.FindUserById(request)
}

func main() {
	lambda.Start(handler)
}
