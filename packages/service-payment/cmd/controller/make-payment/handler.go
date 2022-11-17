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
	secretsManager := secrets.New()
	db, _ := db.New()
	customerRepository := repository.NewCustomerRepository(db)
	customerService, _ := service.NewCustomerService(secretsManager, *customerRepository)
	purchaseService := service.NewPurchaseService()
	controller := controller.NewApiGatewayHandler(customerService, purchaseService)
	return controller.MakePayment(request)
}

func main() {
	lambda.Start(handler)
}
