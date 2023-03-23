package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func TestAcceptBooking_Success(t *testing.T) {
	uid := uuid.New().String()
	_, err := data.BookingService.CreateBooking(uid, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, "pending", data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.AcceptBooking(uid)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, booking.BookingStatus("approved"), b.BookingStatus, "should be equal")
}

func TestAcceptBooking_FailAlreadyAccepted(t *testing.T) {
	uid := uuid.New().String()
	_, err := data.BookingService.CreateBooking(uid, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, "pending", data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	b, err := data.BookingService.AcceptBooking(uid)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, booking.BookingStatus("approved"), b.BookingStatus, "should be equal")

	b, err = data.BookingService.AcceptBooking(uid)
	assert.Error(t, err, "booking status is not pending")
}
