package controller

import "github.com/kokiebisu/rental-storage/service-booking/internal/core/port"

func New(service port.BookingService) *ApiGatewayHandler {
	return NewApiGatewayHandler(service)
}
