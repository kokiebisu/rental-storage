package controller

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
)

func New() *ApiGatewayHandler {
	e := service.NewTokenService()
	c := service.NewCryptoService()
	a := service.NewAuthenticationService(e, c)
	return NewApiGatewayHandler(a)
}
