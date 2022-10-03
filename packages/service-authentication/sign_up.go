package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/bcrypt"
)

type UserCreationPayload struct {
	EmailAddress      string `json:"emailAddress"`
	Password   string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type AuthorizationTokenPayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// get email address and password from event argument
	bodyRequest := UserCreationPayload{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 10)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("cannot hash password"), StatusCode: 500}, nil
	}
	
	updatedUser := &UserCreationPayload {
		EmailAddress: bodyRequest.EmailAddress,
		Password: string(hash),
		FirstName: bodyRequest.FirstName,
		LastName: bodyRequest.LastName,
	}

	fmt.Println(updatedUser)

	encodedUpdatedUser, err := json.Marshal(&updatedUser)
	if err != nil {
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
    }
	userEndpoint := fmt.Sprintf("%s/users", os.Getenv("SERVICE_API_ENDPOINT"))
	// // send arguments to user-service rest endpoint and create row entry
	_, err = http.Post(userEndpoint, "application/json", bytes.NewBuffer(encodedUpdatedUser))
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	
	token := GenerateToken(bodyRequest.EmailAddress)
	fmt.Println("TOKEN: ", token)
	// // attach auth token to redis
	
	// sessionEndpoint := fmt.Sprintf("%s/sessions")
	// request := AuthorizationTokenPayload{
	// 	AuthorizationToken: token,
	// }
	// err := json.Unmarshal(&request)
	// http.Post(sessionEndpoint, "application/json", bytes.NewBuffer())
	tokenPayload := &AuthorizationTokenPayload {
		AuthorizationToken: token,
	}
	encoded, err := json.Marshal(tokenPayload)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	//Returning response with authorization token
	return events.APIGatewayProxyResponse{Body: string(encoded), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}