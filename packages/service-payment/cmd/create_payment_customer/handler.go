package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"

	"service-payment/pkg/adapter"
	"service-payment/pkg/domain"
	"service-payment/pkg/port"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := port.CreatePaymentCustomerBody{}
	json.Unmarshal([]byte(request.Body), &body)
	
	secretsManager := adapter.SetupSecretsManager()
	param, err := secretsManager.GetParam()
	if err != nil {
		panic(err)
	}
	stripe.Key = *param.Parameter.Value
	params := &stripe.CustomerParams{
		// Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
		Name: stripe.String(fmt.Sprintf("%s %s", body.FirstName, body.LastName)),
		Email: stripe.String(body.EmailAddress),
	  }
	c, _ := customer.New(params)
	payload := &port.ResponsePayload {
		ProviderId: c.ID,
		ProviderType: "stripe",
	}
	db, err := adapter.SetupDB()
	if err != nil {
		panic("Unable to setup db")
	}
	repository := adapter.SetupCustomerRepository(db)
	pc := &domain.PaymentCustomer{
		UserId: body.UserId,
		CustomerId: payload.ProviderId,
		ProviderType: payload.ProviderType,
	}
	err = repository.SetupTables()
	if err != nil {
		panic("Unable to setup: " + err.Error())
	}
	_, err = repository.Save(pc)
	if err != nil {
		panic("Unable to save payment customer: " + err.Error())
	}

	return adapter.SendResponse(payload)
}

func main() {
	lambda.Start(handler)
}