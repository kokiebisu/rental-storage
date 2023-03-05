package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingRepository interface {
	Save(booking booking.Entity) *customerror.CustomError
	Delete(id string) *customerror.CustomError
	FindById(id string) (booking.Entity, *customerror.CustomError)
	FindManyByUserId(userId string) ([]booking.Entity, *customerror.CustomError)
	FindManyBySpaceId(spaceId string) ([]booking.Entity, *customerror.CustomError)
}
