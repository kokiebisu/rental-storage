package port

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceFactory interface {
	New(title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, feeCurrency amount.CurrencyType, feeAmount int64, feeType string) (space.Entity, *customerror.CustomError)
}
