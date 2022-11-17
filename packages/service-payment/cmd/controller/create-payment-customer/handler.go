package main

import (
	"service-payment/internal/adapter/controller"
	"service-payment/internal/adapter/db"
	"service-payment/internal/adapter/secrets-manager"
	"service-payment/internal/core/service"
	"service-payment/internal/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := db.New()
	if err != nil {
		panic("Unable to setup db")
	}
	repository := repository.NewCustomerRepository(db)
	secretsManager := secrets.New()
	customerService, _ := service.NewCustomerService(secretsManager, *repository)
	purchaseService := service.NewPurchaseService()
	controller := controller.NewApiGatewayHandler(customerService, purchaseService)
	return controller.CreateCustomer(request)
}

func main() {
	lambda.Start(handler)
}
