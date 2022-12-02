package port

import domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"

type ListingService interface {
	FindListingsWithinLatLng(latitude float32, longitude float32, distance int32) ([]domain.ListingDTO, error)
	FindListingById(uid string) (domain.ListingDTO, error)
	CreateListing(lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, title string, feeAmount int32, feeCurrency domain.CurrencyType, feeType domain.RentalFeeType) (string, error)
	RemoveListingById(uid string) error
}
