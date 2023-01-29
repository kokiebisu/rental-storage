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

type CreateUserResponsePayload struct {
	UId string `json:"uid"`
}

type FindUserByEmailResponsePayload struct {
	User domain.UserDTO `json:"user"`
}

type FindUserByIdResponsePayload struct {
	User domain.UserDTO `json:"user"`
}

type RemoveUserByIdResponsePayload struct {
	UId string `json:"uid"`
}

func NewApiGatewayHandler(service port.UserService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateUser(event events.APIGatewayProxyRequest) (CreateUserResponsePayload, *errors.CustomError) {
	body := struct {
		EmailAddresss string `json:"emailAddress"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Password      string `json:"password"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return CreateUserResponsePayload{}, errors.ErrorHandler.CustomError("unable to unmarshal request body", err)
	}

	uid, err := h.service.CreateUser(body.EmailAddresss, body.FirstName, body.LastName, body.Password)
	payload := CreateUserResponsePayload{
		UId: uid,
	}

	return payload, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) FindUserByEmail(event events.APIGatewayProxyRequest) (FindUserByEmailResponsePayload, *errors.CustomError) {
	emailAddress := event.QueryStringParameters["emailAddress"]
	user, err := h.service.FindByEmail(emailAddress)
	payload := FindUserByEmailResponsePayload{
		User: user.ToDTO(),
	}
	return payload, err
}

func (h *ApiGatewayHandler) FindUserById(event events.APIGatewayProxyRequest) (FindUserByIdResponsePayload, *errors.CustomError) {
	uid := event.PathParameters["userId"]
	user, err := h.service.FindById(uid)
	payload := FindUserByIdResponsePayload{
		User: user.ToDTO(),
	}
	return payload, err
}

func (h *ApiGatewayHandler) RemoveUserById(event events.APIGatewayProxyRequest) (RemoveUserByIdResponsePayload, *errors.CustomError) {
	uid := event.PathParameters["userId"]
	err := h.service.RemoveById(uid)
	payload := RemoveUserByIdResponsePayload{
		UId: uid,
	}
	return payload, err
}
