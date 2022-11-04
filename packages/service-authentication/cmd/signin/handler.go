package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get email address and password from event argument
	bodyRequest := SignInArgument{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", os.Getenv("SERVICE_API_ENDPOINT"), bodyRequest.EmailAddress)
	// check if the email address exists in the user db
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	user := &User{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	matched, err := VerifyPassword(user.Password, bodyRequest.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("something went wrong when validating pasword"), StatusCode: 500}, nil
	}
	if !matched {
		return events.APIGatewayProxyResponse{Body: string("provided password is invalid"), StatusCode: 500}, nil
	}
	response := &Payload {
		UId: user.Uid,
	}
	token, err := GenerateJWTToken(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("unable to generate token"), StatusCode: 500}, nil
	}

	return SendResponse(token)
}



func main() {
	lambda.Start(handler)
}