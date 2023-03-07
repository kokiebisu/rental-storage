package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func TestFindBookings_Success(t *testing.T) {
	mockBooking := data.MockBooking
	_, err := data.BookingService.CreateBooking(mockBooking.UId, mockBooking.UserId, mockBooking.SpaceId, mockBooking.ImageUrls, mockBooking.BookingStatus, mockBooking.Description, mockBooking.StartDate, mockBooking.EndDate, mockBooking.CreatedAt, mockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	b, err := data.BookingService.FindBookingById(mockBooking.UId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(b.UId), 0, "should return valid uid where the lenth is greater than 0")
}
