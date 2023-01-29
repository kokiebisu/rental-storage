package controller

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
)

func New() *ApiGatewayHandler {
	service := service.NewEncryptionService()
	return NewApiGatewayHandler(service)
}
