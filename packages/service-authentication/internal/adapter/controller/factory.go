package controller

import "service-authentication/internal/core/port"

func New(service port.EncryptionService) *ApiGatewayHandler {
	return NewApiGatewayHandler(service)
}
