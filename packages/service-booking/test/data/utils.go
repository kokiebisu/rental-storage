package data

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-booking/mocks"
)

var (
	MockBookingRepo *mocks.BookingRepository
	BookingService  port.BookingService
)
