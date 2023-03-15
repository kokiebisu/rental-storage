package test

import (
	"fmt"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/helper"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var stringifyResults = []StringifyResult{
	{
		adapter.FindSpaceByIdResponsePayload{
			Space: data.MockSpaces[0],
		},
		fmt.Sprintf(`{"space":{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"}}`,
			data.MockSpaces[0].UId,
			data.MockSpaces[0].Title,
			data.MockSpaces[0].LenderId,
			data.MockSpaces[0].Location.Address,
			data.MockSpaces[0].Location.City,
			data.MockSpaces[0].Location.Country,
			data.MockSpaces[0].Location.CountryCode,
			data.MockSpaces[0].Location.Phone,
			data.MockSpaces[0].Location.Province,
			data.MockSpaces[0].Location.ProvinceCode,
			data.MockSpaces[0].Location.Zip,
			data.MockSpaces[0].Location.Coordinate.Latitude,
			data.MockSpaces[0].Location.Coordinate.Longitude,
			data.MockSpaces[0].ImageUrls[0],
			data.MockSpaces[0].ImageUrls[1],
			data.MockSpaces[0].Description,
			data.MockSpaces[0].CreatedAt,
			data.MockSpaces[0].UpdatedAt,
		),
	},
	{
		adapter.FindSpacesResponsePayload{
			Spaces: []space.DTO{
				data.MockSpaces[0],
				data.MockSpaces[0],
			},
		},
		fmt.Sprintf(`{"spaces":[{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"},{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"}]}`,
			data.MockSpaces[0].UId,
			data.MockSpaces[0].Title,
			data.MockSpaces[0].LenderId,
			data.MockSpaces[0].Location.Address,
			data.MockSpaces[0].Location.City,
			data.MockSpaces[0].Location.Country,
			data.MockSpaces[0].Location.CountryCode,
			data.MockSpaces[0].Location.Phone,
			data.MockSpaces[0].Location.Province,
			data.MockSpaces[0].Location.ProvinceCode,
			data.MockSpaces[0].Location.Zip,
			data.MockSpaces[0].Location.Coordinate.Latitude,
			data.MockSpaces[0].Location.Coordinate.Longitude,
			data.MockSpaces[0].ImageUrls[0],
			data.MockSpaces[0].ImageUrls[1],
			data.MockSpaces[0].Description,
			data.MockSpaces[0].CreatedAt,
			data.MockSpaces[0].UpdatedAt,
			data.MockSpaces[0].UId,
			data.MockSpaces[0].Title,
			data.MockSpaces[0].LenderId,
			data.MockSpaces[0].Location.Address,
			data.MockSpaces[0].Location.City,
			data.MockSpaces[0].Location.Country,
			data.MockSpaces[0].Location.CountryCode,
			data.MockSpaces[0].Location.Phone,
			data.MockSpaces[0].Location.Province,
			data.MockSpaces[0].Location.ProvinceCode,
			data.MockSpaces[0].Location.Zip,
			data.MockSpaces[0].Location.Coordinate.Latitude,
			data.MockSpaces[0].Location.Coordinate.Longitude,
			data.MockSpaces[0].ImageUrls[0],
			data.MockSpaces[0].ImageUrls[1],
			data.MockSpaces[0].Description,
			data.MockSpaces[0].CreatedAt,
			data.MockSpaces[0].UpdatedAt,
		),
	},
	{
		adapter.AddSpaceResponsePayload{
			UId: data.MockSpaces[0].UId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, data.MockSpaces[0].UId),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
