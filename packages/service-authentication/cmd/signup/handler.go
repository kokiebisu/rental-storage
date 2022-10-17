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


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get email address and password from event argument
	bodyRequest := UserCreationPayload{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	fmt.Println("Password: ", bodyRequest.Password)

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("cannot hash password"), StatusCode: 500}, nil
	}
	
	updatedUser := &UserCreationPayload {
		EmailAddress: bodyRequest.EmailAddress,
		Password: string(hash),
		FirstName: bodyRequest.FirstName,
		LastName: bodyRequest.LastName,
	}

	encodedUpdatedUser, err := json.Marshal(&updatedUser)
	if err != nil {
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
    }
	userEndpoint := fmt.Sprintf("%s/users", os.Getenv("SERVICE_API_ENDPOINT"))
	
	resp, err := http.Post(userEndpoint, "application/json", bytes.NewBuffer(encodedUpdatedUser))
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	if resp.StatusCode == 500 {
		payload := struct{
			Message string `json:"message"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return events.APIGatewayProxyResponse{Body: string("unable to decode"), StatusCode: 500}, nil
		}
		encodedMessage, err := json.Marshal(payload)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}
		return events.APIGatewayProxyResponse{Body: string(encodedMessage), StatusCode: 500}, nil
	}
	response := &Payload{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return events.APIGatewayProxyResponse{Body: string("unable to decode"), StatusCode: 500}, nil
	}
	token, err := GenerateJWTToken(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	
	tokenPayload := &AuthorizationTokenPayload {
		AuthorizationToken: token,
	}

	encoded, err := json.Marshal(tokenPayload)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(encoded), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}