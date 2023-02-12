package space

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
)

type Entity struct {
	UId           string
	Title         string
	LenderId      string
	StreetAddress streetaddress.ValueObject
	Latitude      coordinate.ValueObject
	Longitude     coordinate.ValueObject
	ImageUrls     []string
	Description   string
}

type DTO struct {
	UId           string   `json:"uid"`
	Title         string   `json:"title"`
	LenderId      string   `json:"lenderId"`
	StreetAddress string   `json:"streetAddress"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	ImageUrls     []string `json:"imageUrls"`
	Description   string   `json:"description"`
}

type Raw struct {
	UId           string
	Title         string
	LenderId      string
	StreetAddress string
	Latitude      float64
	Longitude     float64
	ImageUrls     []string
	Description   string
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
		Description:   r.Description,
	}
}

func (e Entity) ToDTO() DTO {
	return DTO{
		UId:           e.UId,
		Title:         e.Title,
		LenderId:      e.LenderId,
		StreetAddress: e.StreetAddress.Value,
		Latitude:      e.Latitude.Value,
		Longitude:     e.Longitude.Value,
		ImageUrls:     e.ImageUrls,
		Description:   e.Description,
	}
}
