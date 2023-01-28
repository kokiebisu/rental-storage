package port

import (
	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingService interface {
	FindListingsWithinLatLng(latitude float32, longitude float32, distance int32) ([]domain.ListingDTO, *errors.CustomError)
	FindListingById(uid string) (domain.ListingDTO, *errors.CustomError)
	CreateListing(lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, title string, feeAmount int32, feeCurrency domain.CurrencyType, feeType domain.RentalFeeType) (string, *errors.CustomError)
	RemoveListingById(uid string) *errors.CustomError
}
