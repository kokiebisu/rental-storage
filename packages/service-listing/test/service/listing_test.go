package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-listing/mocks"
	"github.com/kokiebisu/rental-storage/service-listing/test/data"
)

// CreateListing
func TestCreateListing_Success(t *testing.T) {
	mockRepo := mocks.NewListingRepository(t)
	mockFactory := mocks.NewListingFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockFeeCurrency, data.MockFeeAmount, data.MockFeeType).Return(data.MockListing, nil)
	mockRepo.On("Save", data.MockListing).Return(data.MockUId, nil)

	listingService := service.NewListingService(mockRepo, mockFactory)
	token, err := listingService.CreateListing(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, int32(data.MockFeeAmount), data.MockFeeCurrency, fee.RentalFeeType(data.MockFeeType))
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(token), 0, "should return valid uid where the length is greater than 0")
}

// FindListingsWithinLatLng
func TestFindListingsWithinLatLng_Success(t *testing.T) {
	mockRepo := mocks.NewListingRepository(t)
	mockFactory := mocks.NewListingFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockFeeCurrency, data.MockFeeAmount, data.MockFeeType).Return(data.MockListing, nil)
	mockRepo.On("Save", data.MockListing).Return(data.MockUId, nil)

	listingService := service.NewListingService(mockRepo, mockFactory)
	_, err := listingService.CreateListing(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, int32(data.MockFeeAmount), data.MockFeeCurrency, fee.RentalFeeType(data.MockFeeType))
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByLatLng", data.MockLatitude, data.MockLongitude, data.MockDistance).Return([]listing.Entity{data.MockListing}, nil)

	result, err := listingService.FindListingsWithinLatLng(data.MockLatitude, data.MockLongitude, data.MockDistance)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return listings greater than 0")
	listingDTO, _ := data.MockListing.ToDTO()
	assert.Equal(t, listingDTO, result[0])

}

// FindListingsByUserId
func TestFindListingsByUserId_Success(t *testing.T) {
	mockRepo := mocks.NewListingRepository(t)
	mockFactory := mocks.NewListingFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockFeeCurrency, data.MockFeeAmount, data.MockFeeType).Return(data.MockListing, nil)
	mockRepo.On("Save", data.MockListing).Return(data.MockUId, nil)

	listingService := service.NewListingService(mockRepo, mockFactory)
	_, err := listingService.CreateListing(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, int32(data.MockFeeAmount), data.MockFeeCurrency, fee.RentalFeeType(data.MockFeeType))
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByUserId", data.MockLenderId).Return([]listing.Entity{data.MockListing}, nil)

	result, err := listingService.FindListingsByUserId(data.MockLenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return listings greater than 0")
	listingDTO, _ := data.MockListing.ToDTO()
	assert.Equal(t, listingDTO, result[0])
}

// FindListingById
func TestFindListingById_Success(t *testing.T) {
	mockRepo := mocks.NewListingRepository(t)
	mockFactory := mocks.NewListingFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockFeeCurrency, data.MockFeeAmount, data.MockFeeType).Return(data.MockListing, nil)
	mockRepo.On("Save", data.MockListing).Return(data.MockUId, nil)

	listingService := service.NewListingService(mockRepo, mockFactory)
	_, err := listingService.CreateListing(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, int32(data.MockFeeAmount), data.MockFeeCurrency, fee.RentalFeeType(data.MockFeeType))
	assert.Nil(t, err, "should not throw error")

	mockRepo.On("FindOneById", data.MockUId).Return(data.MockListing, nil)

	result, err := listingService.FindListingById(data.MockUId)
	listingDTO, _ := data.MockListing.ToDTO()
	assert.Equal(t, listingDTO, result, "should return a listing which matches the uid")
	assert.Nil(t, err, "should not throw error")
}

// RemoveListingById
func TestRemoveListingById_Success(t *testing.T) {
	mockRepo := mocks.NewListingRepository(t)
	mockFactory := mocks.NewListingFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockFeeCurrency, data.MockFeeAmount, data.MockFeeType).Return(data.MockListing, nil)
	mockRepo.On("Save", data.MockListing).Return(data.MockUId, nil)
	mockRepo.On("Delete", data.MockUId).Return(data.MockUId, nil)

	listingService := service.NewListingService(mockRepo, mockFactory)
	_, err := listingService.CreateListing(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, int32(data.MockFeeAmount), data.MockFeeCurrency, fee.RentalFeeType(data.MockFeeType))
	assert.Nil(t, err, "should not throw error")

	uid, err := listingService.RemoveListingById(data.MockUId)
	assert.Greater(t, len(uid), 0, "greater than 0")
	assert.Nil(t, err, "should not throw error")
}
