package adapter

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"

	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

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

func NewApiGatewayAdapter(service port.UserService, publisher port.UserPublisher) port.Controller {
	return &ApiGatewayAdapter{
		service,
		publisher,
	}
}

type ApiGatewayAdapter struct {
	service   port.UserService
	publisher port.UserPublisher
}

func (h *ApiGatewayAdapter) CreateUser(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	logger.Info("Event", zap.Any("event", event))
	body := struct {
		EmailAddresss string `json:"emailAddress"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Password      string `json:"password"`
	}{}
	err := json.Unmarshal([]byte(event.(events.APIGatewayProxyRequest).Body), &body)
	if err != nil {
		logger.Error("Error unmarshalling body", zap.Error(err))
		return CreateUserResponsePayload{}, customerror.ErrorHandler.UnmarshalError("body", err)
	}
	uid, err := h.service.CreateUser("", body.EmailAddresss, body.FirstName, body.LastName, body.Password, []item.DTO{}, "", "")
	payload := CreateUserResponsePayload{
		UId: uid,
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err.(*customerror.CustomError)
}

func (h *ApiGatewayAdapter) FindUserByEmail(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	emailAddress := event.(events.APIGatewayProxyRequest).QueryStringParameters["emailAddress"]
	if emailAddress == "" {
		logger.Error("Email Address not extracted properly")
		return FindUserByEmailResponsePayload{}, customerror.ErrorHandler.GetParameterError("emailAddress")
	}
	logger.Info("Email Address", zap.String("emailAddress", emailAddress))
	user, err := h.service.FindByEmail(emailAddress)
	logger.Info("Found User", zap.Any("user", user))
	payload := FindUserByEmailResponsePayload{
		User: user.ToDTO(),
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err

}

func (h *ApiGatewayAdapter) FindUserById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	logger.Info("Event", zap.Any("event", event))
	uid := event.(events.APIGatewayProxyRequest).PathParameters["userId"]
	if uid == "" {
		logger.Error("Email Address not extracted properly")
		return FindUserByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("userId")
	}
	user, err := h.service.FindById(uid)
	payload := FindUserByIdResponsePayload{
		User: user.ToDTO(),
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err
}

func (h *ApiGatewayAdapter) RemoveUserById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	uid := event.(events.APIGatewayProxyRequest).PathParameters["userId"]
	if uid == "" {
		logger.Error("uid not extracted properly")
		return RemoveUserByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("userId")
	}
	uid, err := h.service.RemoveById(uid)
	payload := RemoveUserByIdResponsePayload{
		UId: uid,
	}
	return payload, err
}
