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

type SignInArgument struct {
	EmailAddress      string `json:"email"`
	Password   string `json:"password"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	
	// get email address and password from event argument
	arg := request.Body
	argument := SignInArgument{}
	json.Unmarshal([]byte(arg), &argument)
	
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", endpoint, argument.EmailAddress)
	// check if the email address exists in the user db
	fmt.Println(userEndpoint)
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	user := struct {
		Id int `json:"id"`
		Uid string `json:"uid"`
		EmailAddress string `json:"emailAddress"`
		Password string `json:"password"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		CreatedAt string `json:"createdAt"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	 
	matched, err := verifyPassword(argument.EmailAddress, user.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("something went wrong when validating pasword"), StatusCode: 500}, nil
	}
	if !matched {
		return events.APIGatewayProxyResponse{Body: string("provided password is invalid"), StatusCode: 500}, nil
	}

	payload := &Payload {
		UId: user.Uid,
	}
	token, err := GenerateJWTToken(payload)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("unable to generate token"), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(token), StatusCode: 200}, nil
}

func verifyPassword(providedPassword string, userPassword string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(providedPassword), 10)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(userPassword))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func main() {
	lambda.Start(handler)
}