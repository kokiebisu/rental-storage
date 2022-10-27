package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

type CreatePaymentCustomerBody struct {
	EmailAddress string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("hello")
	body := CreatePaymentCustomerBody{}
	json.Unmarshal([]byte(request.Body), &body)
	
	secretsManager := initialize()
	param, err := secretsManager.getParam()
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
	payload := &ResponsePayload {
		ProviderId: c.ID,
		ProviderType: "stripe",
	}
	return SendResponse(payload)
}

func main() {
	lambda.Start(handler)
}