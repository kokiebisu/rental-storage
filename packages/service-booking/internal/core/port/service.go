package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
)

type BookingService interface {
	CreateBooking(amountDTO amount.DTO, ownerId string, listingId string, itemsDTO []item.DTO) (string, error)
	FindUserBookings(userId string) ([]booking.Entity, error)
}
