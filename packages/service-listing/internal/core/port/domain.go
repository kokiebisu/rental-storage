package port

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingFactory interface {
	New(title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, feeCurrency amount.CurrencyType, feeAmount int64, feeType string) (listing.Entity, *customerror.CustomError)
}
