package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/bcrypt"
)

type Argument struct {
	EmailAddress      string `json:"email"`
	Password   string `json:"password"`
}

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(email string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Hash to store:", string(hash))

    return base64.StdEncoding.EncodeToString(hash)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	
	// get email address and password from event argument
	arg := request.Body
	argument := Argument{}
	json.Unmarshal([]byte(arg), &argument)
	
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", endpoint, argument.EmailAddress)
	// check if the email address exists in the user db
	fmt.Println(userEndpoint)
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("request to user endpoint failed"), StatusCode: 500}, nil
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("not able to parse body"), StatusCode: 500}, nil
	}

	fmt.Println("BODY: ", body)

	// hash, err := bcrypt.GenerateFromPassword([]byte(argument.Password), 10)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{Body: string("cannot hash password"), StatusCode: 500}, nil
	// }
	// err = bcrypt.CompareHashAndPassword(hash, body.password)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{Body: string("provided password is invalid"), StatusCode: 500}, nil
	// }

	// token := GenerateToken(argument.EmailAddress)
	// send sqs (service session)

	//Returning response with authorization token
	return events.APIGatewayProxyResponse{Body: string("hello"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}