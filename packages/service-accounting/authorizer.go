package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
    AuthorizationToken string                                `json:"authorizationToken"`
	RequestContext     events.AppSyncLambdaAuthorizerRequestContext `json:"requestContext"`
}

func HandleRequest(ctx context.Context, event Event) (*events.AppSyncLambdaAuthorizerResponse, error) {
		token := event.AuthorizationToken
		// check if token exists in redis session (rest api)
		if token == "Authorized" {
			return &events.AppSyncLambdaAuthorizerResponse{
				IsAuthorized: true,
			}, nil
		}
        return &events.AppSyncLambdaAuthorizerResponse{
			IsAuthorized: false,
		}, nil
}

func main() {
        lambda.Start(HandleRequest)
}