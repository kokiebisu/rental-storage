package listing

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/coordinate"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/street_address"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Entity struct {
	Uid           string
	Title         string
	LenderId      string
	StreetAddress streetaddress.ValueObject
	Latitude      coordinate.ValueObject
	Longitude     coordinate.ValueObject
	ImageUrls     []string
	Fee           fee.ValueObject
}

type DTO struct {
	Uid           string   `json:"uid"`
	Title         string   `json:"title"`
	LenderId      string   `json:"lenderId"`
	StreetAddress string   `json:"streetAddress"`
	Latitude      float32  `json:"latitude"`
	Longitude     float32  `json:"longitude"`
	ImageUrls     []string `json:"imageUrls"`
	Fee           fee.DTO  `json:"fee"`
}

type Raw struct {
	Uid           string
	Title         string
	LenderId      string
	StreetAddress string
	Latitude      float32
	Longitude     float32
	ImageUrls     []string
	Fee           fee.Raw
}

func (r Raw) ToEntity() (Entity, *errors.CustomError) {
	validatedLatitude, err := coordinateFactory.New(r.Latitude)
	if err != nil {
		return Entity{}, err
	}
	validatedLongitude, err := coordinateFactory.New(r.Longitude)
	if err != nil {
		return Entity{}, err
	}
	validatedFee, err := feeFactory.New(r.Fee.Amount.Currency, r.Fee.Amount.Value, r.Fee.Type)
	if err != nil {
		return Entity{}, err
	}
	validatedStreetAddress, err := streetAddressFactory.New(r.StreetAddress)
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		Uid:           r.Uid,
		Title:         r.Title,
		LenderId:      r.LenderId,
		StreetAddress: validatedStreetAddress,
		Latitude:      validatedLatitude,
		Longitude:     validatedLongitude,
		ImageUrls:     r.ImageUrls,
		Fee:           validatedFee,
	}, nil
}

func (e Entity) ToDTO() (DTO, *errors.CustomError) {
	return DTO{
		Uid:           e.Uid,
		Title:         e.Title,
		LenderId:      e.LenderId,
		StreetAddress: e.StreetAddress.Value,
		Latitude:      e.Latitude.Value,
		Longitude:     e.Longitude.Value,
		ImageUrls:     e.ImageUrls,
		Fee: fee.DTO{
			Amount: amount.DTO{
				Value:    e.Fee.Amount.Value,
				Currency: string(e.Fee.Amount.Currency),
			},
			Type: string(e.Fee.Type),
		},
	}, nil
}
