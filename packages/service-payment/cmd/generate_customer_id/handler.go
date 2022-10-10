package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	emailAddress := request.QueryStringParameters["emailAddress"]
	firstName := request.QueryStringParameters["firstName"]
	lastName := request.QueryStringParameters["lastName"]
	
	secretsManager := initialize()
	param, err := secretsManager.getParam()
	if err != nil {
		panic(err)
	}
	stripe.Key = *param.Parameter.Value
	params := &stripe.CustomerParams{
		// Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
		Name: stripe.String(fmt.Sprintf("%s %s", firstName, lastName)),
		Email: stripe.String(emailAddress),
	  }
	c, _ := customer.New(params)
	payload := &ResponsePayload {
		CustomerId: c.ID,
		ProviderType: "stripe",
	}
	return SendResponse(payload)
}

func main() {
	lambda.Start(handler)
}