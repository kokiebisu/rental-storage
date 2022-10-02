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

type User struct {
	Id int `json:"id"`
	Uid string `json:"uid"`
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
}



func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	user := User{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	 
	hash, err := bcrypt.GenerateFromPassword([]byte(argument.Password), 10)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("cannot hash password"), StatusCode: 500}, nil
	}
	fmt.Print("string hash: ", string(hash))
	fmt.Print("hash: ", hash)

	fmt.Print("user password: ", []byte(user.Password))
	err = bcrypt.CompareHashAndPassword(hash, []byte(user.Password))
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("provided password is invalid"), StatusCode: 500}, nil
	}
	fmt.Println("matched!")
	token := GenerateToken(argument.EmailAddress)
	fmt.Println("TOKEN: ", token)
	// send sqs (service session)
	sessionEndpoint := fmt.Sprintf("%s/sessions")
	resp, err = http.Post(sessionEndpoint)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("failed to attach session"), StatusCode: 500}, nil
	}

	//Returning response with authorization token
	return events.APIGatewayProxyResponse{Body: string(token), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}