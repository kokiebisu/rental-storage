package port

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingService interface {
	FindListingsWithinLatLng(latitude float64, longitude float64, distance int32) ([]listing.DTO, *errors.CustomError)
	FindListingById(uid string) (listing.DTO, *errors.CustomError)
	CreateListing(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *errors.CustomError)
	RemoveListingById(uid string) *errors.CustomError
}
