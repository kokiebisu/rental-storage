package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ApiGatewayHandler struct {
	service port.UserService
}

type CreateUserResponsePayload struct {
	UId string `json:"uid"`
}

type FindUserByEmailResponsePayload struct {
	User user.DTO `json:"user"`
}

type FindUserByIdResponsePayload struct {
	User user.DTO `json:"user"`
}

type RemoveUserByIdResponsePayload struct {
	UId string `json:"uid"`
}

func (h *ApiGatewayHandler) CreateUser(event events.APIGatewayProxyRequest) (CreateUserResponsePayload, *customerror.CustomError) {
	body := struct {
		EmailAddresss string `json:"emailAddress"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Password      string `json:"password"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return CreateUserResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}

	uid, err := h.service.CreateUser("", body.EmailAddresss, body.FirstName, body.LastName, body.Password, []item.DTO{}, "")
	payload := CreateUserResponsePayload{
		UId: uid,
	}

	return payload, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) FindUserByEmail(event events.APIGatewayProxyRequest) (FindUserByEmailResponsePayload, *customerror.CustomError) {
	emailAddress := event.QueryStringParameters["emailAddress"]
	if emailAddress == "" {
		return FindUserByEmailResponsePayload{}, customerror.ErrorHandler.GetParameterError("emailAddress")
	}
	user, err := h.service.FindByEmail(emailAddress)
	payload := FindUserByEmailResponsePayload{
		User: user.ToDTO(),
	}
	return payload, err
}

func (h *ApiGatewayHandler) FindUserById(event events.APIGatewayProxyRequest) (FindUserByIdResponsePayload, *customerror.CustomError) {
	uid := event.PathParameters["userId"]
	if uid == "" {
		return FindUserByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("userId")
	}
	user, err := h.service.FindById(uid)
	payload := FindUserByIdResponsePayload{
		User: user.ToDTO(),
	}
	return payload, err
}

func (h *ApiGatewayHandler) RemoveUserById(event events.APIGatewayProxyRequest) (RemoveUserByIdResponsePayload, *customerror.CustomError) {
	uid := event.PathParameters["userId"]
	if uid == "" {
		return RemoveUserByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("userId")
	}
	uid, err := h.service.RemoveById(uid)
	payload := RemoveUserByIdResponsePayload{
		UId: uid,
	}
	return payload, err
}
