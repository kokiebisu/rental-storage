package space

import (
	"time"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

func New(uid string, title string, lenderId string, location location.DTO, imageUrls []string, description string, createdAtString string, updatedAtString string) (Entity, *customerror.CustomError) {
	locationEntity := location.ToValueObject()
	var createdAt time.Time
	if createdAtString == "" {
		createdAt = time.Now()
	} else {
		validatedCreatedAt, rawerr := time.Parse(layoutISO, createdAtString)
		if rawerr != nil {
			return Entity{}, customerror.ErrorHandler.ConvertError("createdAt", "time", nil)
		}
		createdAt = validatedCreatedAt
	}
	var updatedAt time.Time
	if updatedAtString == "" {
		updatedAt = time.Now()
	} else {
		validatedUpdatedAt, rawerr := time.Parse(layoutISO, updatedAtString)
		if rawerr != nil {
			return Entity{}, customerror.ErrorHandler.ConvertError("createdAt", "time", nil)
		}
		updatedAt = validatedUpdatedAt
	}
	urls := []ImageUrl{}
	for _, url := range imageUrls {
		if err := ImageUrl(url).validate(); err != nil {
			return Entity{}, err
		}
		urls = append(urls, ImageUrl(url))
	}
	return Entity{
		UId:         uid,
		Title:       title,
		LenderId:    lenderId,
		Location:    locationEntity,
		ImageUrls:   urls,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (u ImageUrl) validate() *customerror.CustomError {
	if u == "" {
		return customerror.ErrorHandler.InvalidValueError("image url", "cannot be empty")
	}
	return nil
}
