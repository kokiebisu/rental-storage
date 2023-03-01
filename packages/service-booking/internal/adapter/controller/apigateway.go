package controller

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
)

type ApiGatewayHandler struct {
	service port.BookingService
}

func NewApiGatewayHandler(service port.BookingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}
