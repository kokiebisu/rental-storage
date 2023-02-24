package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	location "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
)

var (
	MockSpace = space.DTO{
		UId:         faker.UUIDDigit(),
		Title:       "A beautiful garage",
		Description: "BLAH BLAH",
		LenderId:    faker.UUIDDigit(),
		Location: location.DTO{
			Address:      "Blah Blah",
			City:         "Vancouver",
			Country:      "Canada",
			CountryCode:  "CA",
			Phone:        "123-456-7891",
			Province:     "British Columbia",
			ProvinceCode: "BC",
			Zip:          "V193XA",
			Coordinate: coordinate.DTO{
				Latitude:  faker.Latitude(),
				Longitude: faker.Longitude(),
			},
		},
		ImageUrls: []string{
			faker.URL(),
			faker.URL(),
		},
		CreatedAt: faker.Date(),
		UpdatedAt: faker.Date(),
	}
)
