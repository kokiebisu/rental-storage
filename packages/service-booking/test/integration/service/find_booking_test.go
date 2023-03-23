package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func TestFindBookings_Success(t *testing.T) {
	uid := uuid.New().String()
	_, err := data.BookingService.CreateBooking(uid, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, data.MockBooking.BookingStatus, data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	b, err := data.BookingService.FindBookingById(uid)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(b.UId), 0, "should return valid uid where the lenth is greater than 0")
}
