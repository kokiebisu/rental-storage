package port

import (
	"github.com/aws/aws-lambda-go/events"
	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
)

type Controller interface {
	Authorize(event interface{}) (interface{}, *customerror.CustomError)
}

type Event struct {
	AuthorizationToken string                                       `json:"authorizationToken"`
	RequestContext     events.AppSyncLambdaAuthorizerRequestContext `json:"requestContext"`
}
