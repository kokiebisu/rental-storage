package main

import (
	"fmt"
	"os"

	"service-payment/pkg/port"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event port.Event) (string, error) {
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