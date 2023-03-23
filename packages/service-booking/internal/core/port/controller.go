package port

import customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"

type Controller interface {
	CreateBooking(req interface{}) (interface{}, *customerror.CustomError)
	FindBookingById(req interface{}) (interface{}, *customerror.CustomError)
	FindBookings(req interface{}) (interface{}, *customerror.CustomError)
	AcceptBooking(req interface{}) (interface{}, *customerror.CustomError)
}
