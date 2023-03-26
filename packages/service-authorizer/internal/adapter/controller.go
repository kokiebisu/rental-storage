package adapter

import (
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/client"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
	"go.uber.org/zap"
)

func NewApiGatewayAdapter(service port.AuthorizerService) port.Controller {
	return &ApiGatewayAdapter{
		service,
	}
}

type AuthorizeResponsePayload struct {
	UId string `json:"uid"`
	Exp int    `json:"exp"`
}

type ApiGatewayAdapter struct {
	service port.AuthorizerService
}

func (a *ApiGatewayAdapter) Authorize(event interface{}) (interface{}, *customerror.CustomError) {
	logger, err := client.GetLoggerClient()
	if err != nil {
		return AuthorizeResponsePayload{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	logger.Info("Event", zap.Any("event", event))
	token := event.(port.Event).AuthorizationToken

	claim, err := a.service.Authorize(token)
	if err != nil {
		return AuthorizeResponsePayload{}, err
	}
	payload := AuthorizeResponsePayload{
		UId: claim.UId,
		Exp: claim.Exp,
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, nil
}
