package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kokiebisu/rental-storage/service-authorizer/cmd/controller"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/port"
)

func HandleRequest(ctx context.Context, event port.Event) (*events.AppSyncLambdaAuthorizerResponse, error) {
	c, err := controller.New()
	if err != nil {
		return controller.SendUnAuthorizedResponse(), err
	}
	payload, err := c.Authorize(event)
	if err != nil {
		return controller.SendUnAuthorizedResponse(), err
	}
	return controller.SendAuthorizedResponse(payload.(adapter.AuthorizeResponsePayload).UId), nil
}

func main() {
	lambda.Start(HandleRequest)
}
