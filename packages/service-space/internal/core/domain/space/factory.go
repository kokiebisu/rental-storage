package space

import (
	"github.com/google/uuid"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

var (
	coordinateFactory    = &coordinate.Factory{}
	streetAddressFactory = &streetaddress.Factory{}
)

type Factory struct{}

func (f *Factory) New(title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string) (Entity, *customerror.CustomError) {

	validatedLatitude, err := coordinateFactory.New(latitude)
	if err != nil {
		return Entity{}, err
	}
	validatedLongitude, err := coordinateFactory.New(longitude)
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
	}, nil
}
