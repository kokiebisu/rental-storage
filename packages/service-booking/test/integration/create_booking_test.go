package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking_Success(t *testing.T) {
	mockBooking := data.MockBooking
	id, err := data.BookingService.CreateBooking(mockBooking.UId, mockBooking.UserId, mockBooking.SpaceId, mockBooking.ImageUrls, mockBooking.Description, mockBooking.StartDate, mockBooking.EndDate, mockBooking.CreatedAt, mockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
}
