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

// AppSyncLambdaAuthorizerRequestContext contains the parameters of the AppSync invocation which triggered
// this authorization request.
type AppSyncLambdaAuthorizerRequestContext struct {
	APIID         string                 `json:"apiId"`
	AccountID     string                 `json:"accountId"`
	RequestID     string                 `json:"requestId"`
	QueryString   string                 `json:"queryString"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// AppSyncLambdaAuthorizerResponse represents the expected format of an authorization response to AppSync.
type AppSyncLambdaAuthorizerResponse struct {
	IsAuthorized    bool                   `json:"isAuthorized"`
	ResolverContext map[string]interface{} `json:"resolverContext,omitempty"`
	DeniedFields    []string               `json:"deniedFields,omitempty"`
	TTLOverride     *int                   `json:"ttlOverride,omitempty"`
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