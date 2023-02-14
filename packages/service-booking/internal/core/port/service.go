package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService interface {
	CreateBooking(id string, userId string, spaceId string, imageUrls []string, startDate string, endDate string, createdAt string, updatedAt string) (string, *customerror.CustomError)
	FindUserBookings(userId string) ([]booking.Entity, *customerror.CustomError)
	FindById(uid string) (booking.Entity, *customerror.CustomError)
	FindManyBySpaceId(spaceId string) ([]booking.Entity, *customerror.CustomError)
}
