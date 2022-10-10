package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Identity Identity
}

type Identity struct {
	ResolverContext ResolverContext
}

type ResolverContext struct {
	UId string `json:"uid"`
}

func handler(event Event) (string, error) {
	fmt.Println("EVENT: ", event)
	fmt.Println("EVENT Identity: ", event.Identity)
	fmt.Println("EVENT Resolver conext: ", event.Identity.ResolverContext)
	fmt.Println("EVENT Resolver conext uid: ", event.Identity.ResolverContext.UId)

	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	fmt.Println("ENDPOINT: ", endpoint)

	return "HEllo world", nil
}

func main() {
	lambda.Start(handler)
}