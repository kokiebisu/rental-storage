package test

import (
	"testing"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/mocks"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, *customerror.CustomError) {
	data.MockBookingRepo = mocks.NewBookingRepository(t)
	data.MockBookingRepo.On("Save", data.MockBooking.ToEntity()).Return(nil)

	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	token, err := data.BookingService.CreateBooking(data.MockBooking.UId, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
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
	data.MockBookingRepo.On("FindOneById", uid).Return(data.MockBooking.ToEntity(), nil)
	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.FindBookingById(uid)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, b.UId, b.UId)
}

// Find Bookings with status PENDING
func TestFindPendingBookings_Success(t *testing.T) {
	_, err := setupTest(t)
	expected := []booking.Entity{data.MockBooking.ToEntity()}
	data.MockBookingRepo.On("FindManyBySpaceId", data.MockBooking.SpaceId, "PENDING").Return(expected, nil)
	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.FindBookingsBySpaceId(data.MockBooking.SpaceId, booking.PENDING)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, b[0].UId, expected[0].UId)
}

// Find Bookings with status APPROVED
func TestFindApprovedBookings_Success(t *testing.T) {
	_, err := setupTest(t)
	expected := []booking.Entity{data.MockBooking.ToEntity()}
	data.MockBookingRepo.On("FindManyBySpaceId", data.MockBooking.SpaceId, "APPROVED").Return(expected, nil)
	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.FindBookingsBySpaceId(data.MockBooking.SpaceId, booking.APPROVED)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, b[0].UId, expected[0].UId)
}
