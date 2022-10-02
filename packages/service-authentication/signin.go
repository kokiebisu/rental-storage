package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Argument struct {
	EmailAddress      string `json:"email"`
	Password   string `json:"password"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	fmt.Println(endpoint)

	// get email address and password from event argument
	arg := request.Body
	argument := Argument{}
	json.Unmarshal([]byte(arg), &argument)

	// check if the email address exists in the user db
	fmt.Printf("%s, %s", argument.EmailAddress, argument.Password)

	// if exists, bcrypt hash the password and compare
		// if match: generate authorization token and send sqs (service-session)

	//Returning response with authorization token
	return events.APIGatewayProxyResponse{Body: string("hello"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}