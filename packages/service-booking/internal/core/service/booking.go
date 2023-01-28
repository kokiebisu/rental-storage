package service

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingService struct {
	bookingRepository port.BookingRepository
}

func NewBookingService(bookingRepository port.BookingRepository) *BookingService {
	return &BookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *BookingService) CreateBooking(amountDTO amount.DTO, userId string, listingId string, itemsDTO []item.DTO) (string, *errors.CustomError) {
	itemEntities := []item.Entity{}
	amountEntity, err := amount.New(amountDTO.Value, amountDTO.Currency)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	for _, i := range itemsDTO {
		validItem, err := item.New(i.Name, i.ImageUrls)
		if err != nil {
			return "", errors.ErrorHandler.InternalServerError()
		}
		itemEntities = append(itemEntities, validItem)
	}
	bookingEntity, err := booking.New(amountEntity, userId, listingId, itemEntities)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	err = s.bookingRepository.Save(bookingEntity)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	return bookingEntity.Id, nil
}

func (s *BookingService) FindUserBookings(userId string) ([]booking.Entity, *errors.CustomError) {
	bookings, err := s.bookingRepository.FindManyByUserId(userId)
	if err != nil {
		return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	return bookings, nil
}
