package adapter

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type ApiGatewayHandler struct {
	service port.UserService
}

func NewApiGatewayHandler(service port.UserService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := struct {
		EmailAddresss string `json:"emailAddress"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Password      string `json:"password"`
	}{}
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		panic(err)
	}
	err = h.service.CreateUser(body.EmailAddresss, body.FirstName, body.LastName, body.Password)
	if err != nil {
		panic(err)
	}
	return sendCreatedResponse()
}

func (h *ApiGatewayHandler) FindUserByEmail(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	emailAddress := request.QueryStringParameters["emailAddress"]
	user, err := h.service.FindByEmail(emailAddress)
	if err != nil {
		panic(err)
	}
	return sendResponse(user.ToDTO())
}

func (h *ApiGatewayHandler) FindUserById(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uid := request.PathParameters["userId"]
	user, err := h.service.FindById(uid)
	if err != nil {
		panic(err)
	}
	return sendResponse(user.ToDTO())
}

func (h *ApiGatewayHandler) RemoveUserById(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uid := request.PathParameters["userId"]
	err := h.service.RemoveById(uid)
	if err != nil {
		panic(err)
	}
	return sendDeletedResponse()
}

func sendResponse(data interface{}) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(encoded),
		StatusCode: 200,
	}, nil
}

func sendDeletedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func sendCreatedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}
