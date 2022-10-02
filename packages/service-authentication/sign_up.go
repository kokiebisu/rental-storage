package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/bcrypt"
)

type SignUpArgument struct {
	EmailAddress      string `json:"email"`
	Password   string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// get email address and password from event argument
	arg := request.Body
	argument := SignUpArgument{}
	json.Unmarshal([]byte(arg), &argument)
	 
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(argument.Password), 10)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("cannot hash password"), StatusCode: 500}, nil
	}
	
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	userEndpoint := fmt.Sprintf("%s/users", endpoint)
	// send arguments to user-service rest endpoint and create row entry
	http.Post(userEndpoint)
	token := GenerateToken(argument.EmailAddress)
	fmt.Println("TOKEN: ", token)
	// attach auth token to redis
	sessionEndpoint := fmt.Sprintf("%s/sessions")
	http.Post(sessionEndpoint)

	//Returning response with authorization token
	return events.APIGatewayProxyResponse{Body: string(token), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}