package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type ApiGatewayHandler struct {
	service port.EncryptionService
}

type SignInResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type SignUpResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type VerifyResponsePayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

func NewApiGatewayHandler(service port.EncryptionService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) SignIn(event events.APIGatewayProxyRequest) (SignInResponsePayload, *errors.CustomError) {

	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{}

	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return SignInResponsePayload{}, errors.ErrorHandler.BodyRequestParseError()
	}

	token, err := h.service.SignIn(bodyRequest.EmailAddress, bodyRequest.Password)
	payload := SignInResponsePayload{
		AuthorizationToken: token,
	}
	return payload, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) SignUp(event events.APIGatewayProxyRequest) (SignUpResponsePayload, *errors.CustomError) {
	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return SignUpResponsePayload{}, errors.ErrorHandler.BodyRequestParseError()
	}
	token, err := h.service.SignUp(bodyRequest.EmailAddress, bodyRequest.FirstName, bodyRequest.LastName, bodyRequest.Password)
	payload := SignUpResponsePayload{
		AuthorizationToken: token,
	}
	return payload, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) Verify(event events.APIGatewayProxyRequest) (VerifyResponsePayload, *errors.CustomError) {
	bodyRequest := struct {
		AuthorizationToken string `json:"authorizationToken"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return VerifyResponsePayload{}, errors.ErrorHandler.BodyRequestParseError()
	}
	token, err := h.service.Verify(bodyRequest.AuthorizationToken)
	payload := VerifyResponsePayload{
		AuthorizationToken: token,
	}
	return payload, err.(*errors.CustomError)
}
