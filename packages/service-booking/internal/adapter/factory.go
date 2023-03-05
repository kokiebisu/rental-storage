package adapter

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

func NewControllerAdapter(service port.BookingService) (port.Controller, *customerror.CustomError) {
	return NewApiGatewayAdapter(service)
}
