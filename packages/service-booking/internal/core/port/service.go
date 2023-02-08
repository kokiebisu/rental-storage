package port

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService interface {
	CreateBooking(id string, amountDTO amount.DTO, userId string, listingId string, itemsDTO []item.DTO, createdAt string, updatedAt string) (string, *customerror.CustomError)
	FindUserBookings(userId string) ([]booking.Entity, *customerror.CustomError)
	FindById(uid string) (booking.Entity, *customerror.CustomError)
}
