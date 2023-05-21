package adapter

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/client"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type SignInResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
	RefreshToken       string `json:"refreshToken"`
}

type SignUpResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
	RefreshToken       string `json:"refreshToken"`
}

type VerifyResponsePayload struct {
	UId string `json:"uid"`
	Exp int64  `json:"exp"`
}

type ApiGatewayAdapter struct {
	service port.AuthenticationService
}

func NewApiGatewayAdapter(service port.AuthenticationService) port.Controller {
	return &ApiGatewayAdapter{
		service,
	}
}

func (h *ApiGatewayAdapter) SignIn(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{}

	err := json.Unmarshal([]byte(event.(events.APIGatewayProxyRequest).Body), &bodyRequest)
	if err != nil {
		logger.Error(err.Error())
		return SignInResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}

	token, cerr := h.service.SignIn(bodyRequest.EmailAddress, bodyRequest.Password)
	if cerr != nil {
		logger.Error(cerr.Error())
		return SignInResponsePayload{}, cerr
	}
	payload := SignInResponsePayload{
		AuthorizationToken: string(token["access_token"]),
		RefreshToken:       string(token["refresh_token"]),
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, cerr
}

func (h *ApiGatewayAdapter) SignUp(event interface{}) (interface{}, *customerror.CustomError) {
	// get email address and password from event argument
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{}
	err := json.Unmarshal([]byte(event.(events.APIGatewayProxyRequest).Body), &bodyRequest)
	if err != nil {
		logger.Error(err.Error())
		return SignUpResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}
	token, cerr := h.service.SignUp(bodyRequest.EmailAddress, bodyRequest.FirstName, bodyRequest.LastName, bodyRequest.Password)
	if cerr != nil {
		return SignUpResponsePayload{}, cerr
	}
	payload := SignUpResponsePayload{
		AuthorizationToken: string(token["access_token"]),
		RefreshToken:       string(token["refresh_token"]),
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, nil
}

func (h *ApiGatewayAdapter) Verify(event interface{}) (interface{}, *customerror.CustomError) {
	var token string
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	token = event.(events.APIGatewayProxyRequest).QueryStringParameters["authorizationToken"]
	if token == "" {
		return VerifyResponsePayload{}, customerror.ErrorHandler.QueryParamEmptyError("authorizationToken")
	}
	claims, cerr := h.service.Verify(token)
	if cerr != nil {
		logger.Error(cerr.Error())
		return VerifyResponsePayload{}, cerr
	}
	payload := VerifyResponsePayload{
		UId: claims.UId,
		Exp: claims.ExpiresAt,
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, nil
}
