package space

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type Entity struct {
	UId           string
	Title         string
	LenderId      string
	StreetAddress streetaddress.ValueObject
	Latitude      coordinate.ValueObject
	Longitude     coordinate.ValueObject
	ImageUrls     []string
	Fee           fee.ValueObject
}

type DTO struct {
	UId           string   `json:"uid"`
	Title         string   `json:"title"`
	LenderId      string   `json:"lenderId"`
	StreetAddress string   `json:"streetAddress"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	ImageUrls     []string `json:"imageUrls"`
	Fee           fee.DTO  `json:"fee"`
}

type Raw struct {
	UId           string
	Title         string
	LenderId      string
	StreetAddress string
	Latitude      float64
	Longitude     float64
	ImageUrls     []string
	Fee           fee.Raw
}

func (r Raw) ToEntity() Entity {
	return Entity{
		UId:           r.UId,
		Title:         r.Title,
		LenderId:      r.LenderId,
		StreetAddress: streetaddress.ValueObject{Value: r.StreetAddress},
		Latitude:      coordinate.ValueObject{Value: r.Latitude},
		Longitude:     coordinate.ValueObject{Value: r.Longitude},
		ImageUrls:     r.ImageUrls,
		Fee: fee.ValueObject{
			Amount: amount.ValueObject{
				Value:    r.Fee.Amount.Value,
				Currency: amount.CurrencyType(r.Fee.Amount.Currency),
			},
			Type: fee.RentalFeeType(r.Fee.Type),
		},
	}
}

func (e Entity) ToDTO() (DTO, *customerror.CustomError) {
	return DTO{
		UId:           e.UId,
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
