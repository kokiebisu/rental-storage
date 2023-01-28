package controller

import "github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"

func New(service port.EncryptionService) *ApiGatewayHandler {
	return NewApiGatewayHandler(service)
}
