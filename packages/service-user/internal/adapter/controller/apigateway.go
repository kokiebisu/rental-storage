package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ApiGatewayHandler struct {
	service port.UserService
}

func NewApiGatewayHandler(service port.UserService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateUser(event events.APIGatewayProxyRequest) (string, *errors.CustomError) {
	body := struct {
		EmailAddresss string `json:"emailAddress"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Password      string `json:"password"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}

	uid, err := h.service.CreateUser(body.EmailAddresss, body.FirstName, body.LastName, body.Password)
	return uid, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) FindUserByEmail(event events.APIGatewayProxyRequest) (domain.UserDTO, *errors.CustomError) {
	emailAddress := event.QueryStringParameters["emailAddress"]
	user, err := h.service.FindByEmail(emailAddress)
	if err != nil {
		return domain.UserDTO{}, errors.ErrorHandler.InternalServerError()
	}
	return user.ToDTO(), nil
}

func (h *ApiGatewayHandler) FindUserById(event events.APIGatewayProxyRequest) (domain.UserDTO, *errors.CustomError) {
	uid := event.PathParameters["userId"]
	user, err := h.service.FindById(uid)
	if err != nil {
		return domain.UserDTO{}, errors.ErrorHandler.InternalServerError()
	}
	return user.ToDTO(), nil
}

func (h *ApiGatewayHandler) RemoveUserById(event events.APIGatewayProxyRequest) (string, *errors.CustomError) {
	uid := event.PathParameters["userId"]
	err := h.service.RemoveById(uid)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	return uid, nil
}
