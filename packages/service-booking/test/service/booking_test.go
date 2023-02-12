package test

import (
	"testing"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/mocks"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, *customerror.CustomError) {
	data.MockBookingRepo = mocks.NewBookingRepository(t)
	data.MockBookingRepo.On("Save", data.MockBookingEntity).Return(nil)

	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	token, err := data.BookingService.CreateBooking(data.MockUId, data.MockAmount, data.MockUserId, data.MockSpaceId, []item.DTO{data.MockItem}, data.MockDateString, data.MockDateString)
	return token, err
}

// CreateUser
func TestCreateBooking_Success(t *testing.T) {
	uid, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
}

// FindBookingById
func TestFindBookingById_Success(t *testing.T) {
	uid, err := setupTest(t)
	data.MockBookingRepo.On("FindOneById", uid).Return(data.MockBookingEntity, nil)
	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.FindById(uid)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, b.UId, b.UId)
}

// FindBookingById
func TestFindBookings_Success(t *testing.T) {
	_, err := setupTest(t)
	expected := []booking.Entity{data.MockBookingEntity}
	data.MockBookingRepo.On("FindManyBySpaceId", data.MockSpaceId).Return(expected, nil)
	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.FindManyBySpaceId(data.MockSpaceId)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, b[0].UId, expected[0].UId)
}
