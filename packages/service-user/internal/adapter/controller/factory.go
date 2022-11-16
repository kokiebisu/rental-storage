package controller

import "github.com/kokiebisu/rental-storage/service-user/internal/core/port"

func New(service port.UserService) *ApiGatewayHandler {
	return NewApiGatewayHandler(service)
}
