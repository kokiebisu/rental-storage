package controller

import (
	"encoding/json"
	"service-authentication/internal/core/port"

	"github.com/aws/aws-lambda-go/events"
)

type ApiGatewayHandler struct {
	service port.EncryptionService
}

func NewApiGatewayHandler(service port.EncryptionService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) SignIn(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{}

	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	token, err := h.service.SignIn(bodyRequest.EmailAddress, bodyRequest.Password)
	if err != nil {
		return sendFailureResponse(err)
	}

	return sendResponse(token)
}

func (h *ApiGatewayHandler) SignUp(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get email address and password from event argument
	bodyRequest := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	encoded, err := h.service.SignUp(bodyRequest.EmailAddress, bodyRequest.FirstName, bodyRequest.LastName, bodyRequest.Password)
	return events.APIGatewayProxyResponse{Body: string(encoded), StatusCode: 200}, nil
}

func (h *ApiGatewayHandler) Verify(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bodyRequest := struct {
		AuthorizationToken string `json:"authorizationToken"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &bodyRequest)
	encoded, err := h.service.Verify(bodyRequest.AuthorizationToken)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(encoded), StatusCode: 200}, nil
}

func sendFailureResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       string(err.Error()),
	}, nil
}

func sendResponse(token string) (events.APIGatewayProxyResponse, error) {
	payload := &struct {
		AuthorizationToken string `json:"authorizationToken"`
	}{
		AuthorizationToken: token,
	}
	encoded, _ := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		Body:       string(encoded),
		StatusCode: 200,
	}, nil
}
