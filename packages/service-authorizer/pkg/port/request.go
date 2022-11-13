package port

import "github.com/aws/aws-lambda-go/events"

type Event struct {
    AuthorizationToken string                                `json:"authorizationToken"`
	RequestContext     events.AppSyncLambdaAuthorizerRequestContext `json:"requestContext"`
}