package space

import (
	"github.com/google/uuid"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

var (
	coordinateFactory    = &coordinate.Factory{}
	feeFactory           = &fee.Factory{}
	streetAddressFactory = &streetaddress.Factory{}
)

type Factory struct{}

func (f *Factory) New(title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, feeCurrency amount.CurrencyType, feeAmount int64, feeType string) (Entity, *customerror.CustomError) {

	validatedLatitude, err := coordinateFactory.New(latitude)
	if err != nil {
		return Entity{}, err
	}
	validatedLongitude, err := coordinateFactory.New(longitude)
	if err != nil {
		return Entity{}, err
	}
	validatedFee, err := feeFactory.New(string(feeCurrency), feeAmount, feeType)
	if err != nil {
		return Entity{}, err
	}
	validatedStreetAddress, err := streetAddressFactory.New(streetAddress)
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		UId:           uuid.New().String(),
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: validatedStreetAddress,
		Latitude:      validatedLatitude,
		Longitude:     validatedLongitude,
		ImageUrls:     imageUrls,
		Fee:           validatedFee,
	}, nil
}
