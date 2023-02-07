package test

import (
	"testing"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/mocks"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-booking/test/data"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, *customerror.CustomError) {
	data.MockBookingRepo = mocks.NewBookingRepository(t)
	data.MockBookingRepo.On("Save", data.MockBookingEntity).Return(nil)

	data.BookingService = service.NewBookingService(data.MockBookingRepo)
	token, err := data.BookingService.CreateBooking(data.MockUId, data.MockAmount, data.MockUserId, data.MockListingId, []item.DTO{data.MockItem}, data.MockDateString, data.MockDateString)
	return token, err
}

// CreateUser
func TestCreateBooking_Success(t *testing.T) {
	uid, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
}
