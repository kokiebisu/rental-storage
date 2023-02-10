package port

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingService interface {
	FindListingsWithinLatLng(latitude float64, longitude float64, distance int32) ([]listing.DTO, *customerror.CustomError)
	FindListingById(uid string) (listing.DTO, *customerror.CustomError)
	FindListingsByUserId(userId string) ([]listing.DTO, *customerror.CustomError)
	CreateListing(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *customerror.CustomError)
	RemoveListingById(uid string) (string, *customerror.CustomError)
}
