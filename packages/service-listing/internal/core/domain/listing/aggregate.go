package domain

import (
	"github.com/google/uuid"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Listing struct {
	Uid           string
	Title         string
	LenderId      string
	StreetAddress StreetAddress
	Latitude      Coordinate
	Longitude     Coordinate
	ImageUrls     []string
	Fee           Fee
}

type ListingDTO struct {
	Uid           string   `json:"uid"`
	Title         string   `json:"title"`
	LenderId      string   `json:"lenderId"`
	StreetAddress string   `json:"streetAddress"`
	Latitude      float32  `json:"latitude"`
	Longitude     float32  `json:"longitude"`
	ImageUrls     []string `json:"imageUrls"`
	Fee           FeeDTO   `json:"fee"`
}

type ListingRaw struct {
	Uid           string
	Title         string
	LenderId      string
	StreetAddress string
	Latitude      float32
	Longitude     float32
	ImageUrls     []string
	Fee           FeeRaw
}

func NewListing(title string, lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, feeCurrency CurrencyType, feeAmount int64, feeType string) (Listing, *errors.CustomError) {
	validatedLatitude, err := NewCoordinate(latitude)
	if err != nil {
		return Listing{}, err
	}
	validatedLongitude, err := NewCoordinate(longitude)
	if err != nil {
		return Listing{}, err
	}
	validatedFee, err := NewFee(string(feeCurrency), feeAmount, feeType)
	if err != nil {
		return Listing{}, err
	}
	validatedStreetAddress, err := NewStreetAddress(streetAddress)
	if err != nil {
		return Listing{}, err
	}
	return Listing{
		Uid:           uuid.New().String(),
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: validatedStreetAddress,
		Latitude:      validatedLatitude,
		Longitude:     validatedLongitude,
		ImageUrls:     imageUrls,
		Fee:           validatedFee,
	}, nil
}

func (r ListingRaw) ToEntity() (Listing, *errors.CustomError) {
	validatedLatitude, err := NewCoordinate(r.Latitude)
	if err != nil {
		return Listing{}, err
	}
	validatedLongitude, err := NewCoordinate(r.Longitude)
	if err != nil {
		return Listing{}, err
	}
	validatedFee, err := NewFee(r.Fee.Amount.Currency, r.Fee.Amount.Value, r.Fee.Type)
	if err != nil {
		return Listing{}, err
	}
	validatedStreetAddress, err := NewStreetAddress(r.StreetAddress)
	if err != nil {
		return Listing{}, err
	}
	return Listing{
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

func (e Listing) ToDTO() (ListingDTO, *errors.CustomError) {
	return ListingDTO{
		Uid:           e.Uid,
		Title:         e.Title,
		LenderId:      e.LenderId,
		StreetAddress: e.StreetAddress.Value,
		Latitude:      e.Latitude.Value,
		Longitude:     e.Longitude.Value,
		ImageUrls:     e.ImageUrls,
		Fee: FeeDTO{
			Amount: AmountDTO{
				Value:    e.Fee.Amount.Value,
				Currency: string(e.Fee.Amount.Currency),
			},
			Type: string(e.Fee.Type),
		},
	}, nil
}
