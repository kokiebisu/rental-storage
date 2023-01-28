package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingRepository interface {
	Save(booking booking.Entity) *errors.CustomError
	Delete(id string) *errors.CustomError
	FindOneById(id string) (booking.Entity, *errors.CustomError)
	FindManyByUserId(userId string) ([]booking.Entity, *errors.CustomError)
}
