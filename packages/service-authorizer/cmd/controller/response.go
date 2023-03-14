package controller

import (
	"github.com/aws/aws-lambda-go/events"
)

func SendUnAuthorizedResponse() *events.AppSyncLambdaAuthorizerResponse {
	return &events.AppSyncLambdaAuthorizerResponse{
		IsAuthorized: false,
	}
}

func SendAuthorizedResponse(uid string) *events.AppSyncLambdaAuthorizerResponse {
	return &events.AppSyncLambdaAuthorizerResponse{
		IsAuthorized: true,
		ResolverContext: map[string]interface{}{
			"uid": uid,
		},
	}
}
