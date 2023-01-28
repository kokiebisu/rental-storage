package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService interface {
	CreateBooking(amountDTO amount.DTO, userId string, listingId string, itemsDTO []item.DTO) (string, *errors.CustomError)
	FindUserBookings(userId string) ([]booking.Entity, *errors.CustomError)
}
