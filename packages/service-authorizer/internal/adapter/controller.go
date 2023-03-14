package adapter

import (
	"encoding/json"

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
	var err error
	logger, cerr := client.GetLoggerClient()
	if cerr != nil {
		return AuthorizeResponsePayload{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	logger.Info("Event", zap.Any("event", event))
	jwt := event.(port.Event).AuthorizationToken
	body := struct {
		Token string `json:"AuthorizationToken"`
	}{
		Token: jwt,
	}
	encodedPayload, err := json.Marshal(&body)
	if err != nil {
		return AuthorizeResponsePayload{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	claim, cerr := a.service.Authorize(encodedPayload)
	if cerr != nil {
		return AuthorizeResponsePayload{}, cerr
	}
	payload := AuthorizeResponsePayload{
		UId: claim.UId,
		Exp: claim.Exp,
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, nil
}
