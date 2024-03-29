package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService interface {
	CreateBooking(id string, userId string, spaceId string, imageUrls []string, bookingStatus string, description string, startDate string, endDate string, createdAt string, updatedAt string) (string, *customerror.CustomError)
	AcceptBooking(uid string) (booking.Entity, *customerror.CustomError)
	FindBookingById(uid string) (booking.Entity, *customerror.CustomError)
	FindBookingsBySpaceId(spaceId string, bookingStatus string) ([]booking.Entity, *customerror.CustomError)
	FindBookingsByUserId(userId string, bookingStatus string) ([]booking.Entity, *customerror.CustomError)
}
