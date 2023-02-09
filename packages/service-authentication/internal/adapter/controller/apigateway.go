package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type ApiGatewayHandler struct {
	service port.AuthenticationService
}

type SignInResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type SignUpResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type VerifyResponsePayload struct {
	UId string `json:"uid"`
	Exp int64  `json:"exp"`
}

func NewApiGatewayHandler(service port.AuthenticationService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) SignIn(event events.APIGatewayProxyRequest) (SignInResponsePayload, *customerror.CustomError) {

	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{}

	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return SignInResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}

	token, err := h.service.SignIn(bodyRequest.EmailAddress, bodyRequest.Password)
	payload := SignInResponsePayload{
		AuthorizationToken: token,
	}
	return payload, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) SignUp(event events.APIGatewayProxyRequest) (SignUpResponsePayload, *customerror.CustomError) {
	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return SignUpResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}
	token, err := h.service.SignUp(bodyRequest.EmailAddress, bodyRequest.FirstName, bodyRequest.LastName, bodyRequest.Password)
	payload := SignUpResponsePayload{
		AuthorizationToken: token,
	}
	return payload, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) Verify(event events.APIGatewayProxyRequest) (VerifyResponsePayload, *customerror.CustomError) {
	bodyRequest := struct {
		AuthorizationToken string `json:"authorizationToken"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return VerifyResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}
	claims, err := h.service.Verify(bodyRequest.AuthorizationToken)

	payload := VerifyResponsePayload{
		UId: claims.UId,
		Exp: claims.ExpiresAt,
	}
	return payload, err.(*customerror.CustomError)
}
