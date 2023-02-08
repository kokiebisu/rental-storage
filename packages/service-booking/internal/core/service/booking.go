package service

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService struct {
	bookingRepository port.BookingRepository
}

func NewBookingService(bookingRepository port.BookingRepository) *BookingService {
	return &BookingService{
		bookingRepository,
	}
}

func (s *BookingService) CreateBooking(id string, amountDTO amount.DTO, userId string, listingId string, itemsDTO []item.DTO, createdAt string, updatedAt string) (string, *customerror.CustomError) {
	itemEntities := []item.Entity{}
	amountEntity, err := amount.New(amountDTO.Value, amountDTO.Currency)
	if err != nil {
		return "", err
	}
	for _, i := range itemsDTO {
		validItem, err := item.New(i.Id, i.Name, i.ImageUrls)
		if err != nil {
			return "", err
		}
		itemEntities = append(itemEntities, validItem)
	}
	bookingEntity, err := booking.New(id, amountEntity, userId, listingId, itemEntities, createdAt, updatedAt)
	if err != nil {
		return "", err
	}
	err = s.bookingRepository.Save(bookingEntity)
	if err != nil {
		return "", err
	}
	return bookingEntity.Id, nil
}

// @deprecated not used
func (s *BookingService) FindUserBookings(userId string) ([]booking.Entity, *customerror.CustomError) {
	return s.bookingRepository.FindManyByUserId(userId)
}

func (s *BookingService) FindById(uid string) (booking.Entity, *customerror.CustomError) {
	return s.bookingRepository.FindOneById(uid)
}
