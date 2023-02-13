package space

import (
	"time"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

func New(uid string, title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, description string, createdAtString string) (Entity, *customerror.CustomError) {
	validatedLatitude, err := coordinate.New(latitude)
	if err != nil {
		return Entity{}, err
	}
	validatedLongitude, err := coordinate.New(longitude)
	if err != nil {
		return Entity{}, err
	}
	validatedStreetAddress, err := streetaddress.New(streetAddress)
	if err != nil {
		return Entity{}, err
	}
	var createdAt time.Time
	if createdAtString == "" {
		createdAt = time.Now()
	} else {
		validatedCreatedAt, rawerr := time.Parse(layoutISO, createdAtString)
		if rawerr != nil {
			return Entity{}, customerror.ErrorHandler.ConvertError("createdAt", "time", err)
		}
		createdAt = validatedCreatedAt
	}
	return Entity{
		UId:           uid,
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: validatedStreetAddress,
		Latitude:      validatedLatitude,
		Longitude:     validatedLongitude,
		ImageUrls:     imageUrls,
		Description:   description,
		CreatedAt:     createdAt,
	}, nil
}
