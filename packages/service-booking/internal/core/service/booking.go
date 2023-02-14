package service

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
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

func (s *BookingService) CreateBooking(id string, userId string, spaceId string, imageUrls []string, startDate string, endDate string, createdAt string, updatedAt string) (string, *customerror.CustomError) {
	bookingEntity, err := booking.New(id, userId, spaceId, imageUrls, startDate, endDate, createdAt, updatedAt)
	if err != nil {
		return "", err
	}
	err = s.bookingRepository.Save(bookingEntity)
	if err != nil {
		return "", err
	}
	return bookingEntity.UId, nil
}

// @deprecated not used
func (s *BookingService) FindUserBookings(userId string) ([]booking.Entity, *customerror.CustomError) {
	return s.bookingRepository.FindManyByUserId(userId)
}

func (s *BookingService) FindById(uid string) (booking.Entity, *customerror.CustomError) {
	return s.bookingRepository.FindOneById(uid)
}

func (s *BookingService) FindManyBySpaceId(spaceId string) ([]booking.Entity, *customerror.CustomError) {
	return s.bookingRepository.FindManyBySpaceId(spaceId)
}
