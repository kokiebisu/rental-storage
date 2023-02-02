package test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/coordinate"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/street_address"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/service"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
	"github.com/kokiebisu/rental-storage/service-listing/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	uid           = faker.UUIDDigit()
	title         = "A beautiful garage"
	lenderId      = faker.UUIDDigit()
	streetAddress = faker.FirstName()
	latitude      = float32(faker.Latitude())
	longitude     = float32(faker.Longitude())
	distance      = int32(5)
	imageUrls     = []string{
		faker.URL(),
		faker.URL(),
	}
	feeCurrency = amount.CurrencyType("CAD")
	feeAmount   = int64(50)
	feeType     = "MONTHLY"
	l           = listing.Entity{
		Uid:           uid,
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: streetaddress.ValueObject{Value: streetAddress},
		Latitude:      coordinate.ValueObject{Value: latitude},
		Longitude:     coordinate.ValueObject{Value: longitude},
		ImageUrls:     imageUrls,
		Fee: fee.ValueObject{
			Amount: amount.ValueObject{
				Value:    feeAmount,
				Currency: feeCurrency,
			},
			Type: "MONTHLY",
		},
	}
	s           *service.ListingService
	mockRepo    *mocks.ListingRepository
	mockFactory *mocks.ListingFactory
)

func setupTest(t *testing.T) (string, *errors.CustomError) {
	mockRepo = mocks.NewListingRepository(t)
	mockFactory = mocks.NewListingFactory(t)
	mockFactory.On("New", title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, feeAmount, feeType).Return(l, nil)
	mockRepo.On("Save", l).Return(uid, nil)

	s = service.NewListingService(mockRepo, mockFactory)
	token, err := s.CreateListing(lenderId, streetAddress, latitude, longitude, imageUrls, title, int32(feeAmount), feeCurrency, fee.RentalFeeType(feeType))
	return token, err
}

// CreateListing
func TestCreateListingSuccess(t *testing.T) {
	uid, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}

	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
	assert.Nil(t, err, "should not throw error")
}

// FindListingsWithinLatLng
func TestFindListingsWithinLatLngSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}

	mockRepo.On("FindManyByLatLng", latitude, longitude, distance).Return([]listing.Entity{l}, nil)

	result, err := s.FindListingsWithinLatLng(latitude, longitude, distance)
	assert.Greater(t, len(result), 0, "should return listings greater than 0")
	listingDTO, _ := l.ToDTO()
	assert.Equal(t, listingDTO, result[0])
	assert.Nil(t, err, "should not throw error")
}

// FindListingById
func TestFindListingByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}

	mockRepo.On("FindOneById", uid).Return(l, nil)

	result, err := s.FindListingById(uid)
	listingDTO, _ := l.ToDTO()
	assert.Equal(t, listingDTO, result, "should return a listing which matches the uid")
	assert.Nil(t, err, "should not throw error")
}

// RemoveListingById
func TestRemoveListingByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	mockRepo.On("Delete", uid).Return(nil)

	err = s.RemoveListingById(uid)
	assert.Nil(t, err, "should not throw error")
}
