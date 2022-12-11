package port

import "github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"

type BookingRepository interface {
	Save(booking booking.Entity) error
	Delete(id string) error
	FindOneById(id string) (booking.Entity, error)
	FindManyByUserId(userId string) ([]booking.Entity, error)
}
