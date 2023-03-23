package integration

import (
	"testing"

	"github.com/docker/distribution/uuid"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func TestFindBookingsBySpaceId_PendingStatusSuccess(t *testing.T) {
	uid := uuid.Generate().String()
	_, err := data.BookingService.CreateBooking(uid, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, "pending", data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	bs, err := data.BookingService.FindBookingsBySpaceId(data.MockBooking.SpaceId, "pending")
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(bs), 0, "should return valid uid where the lenth is greater than 0")
}

func TestFindBookingsBySpaceId_ApprovedStatusSuccess(t *testing.T) {
	_, err := data.BookingService.CreateBooking(data.MockBooking.UId, data.MockBooking.UserId, data.MockBooking.SpaceId, data.MockBooking.ImageUrls, "approved", data.MockBooking.Description, data.MockBooking.StartDate, data.MockBooking.EndDate, data.MockBooking.CreatedAt, data.MockBooking.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	bs, err := data.BookingService.FindBookingsBySpaceId(data.MockBooking.SpaceId, "approved")
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(bs), 0, "should return valid uid where the lenth is greater than 0")
}
