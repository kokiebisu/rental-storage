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
			Space: data.MockSpace,
		},
		fmt.Sprintf(`{"space":{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"}}`,
			data.MockSpace.UId,
			data.MockSpace.Title,
			data.MockSpace.LenderId,
			data.MockSpace.Location.Address,
			data.MockSpace.Location.City,
			data.MockSpace.Location.Country,
			data.MockSpace.Location.CountryCode,
			data.MockSpace.Location.Phone,
			data.MockSpace.Location.Province,
			data.MockSpace.Location.ProvinceCode,
			data.MockSpace.Location.Zip,
			data.MockSpace.Location.Coordinate.Latitude,
			data.MockSpace.Location.Coordinate.Longitude,
			data.MockSpace.ImageUrls[0],
			data.MockSpace.ImageUrls[1],
			data.MockSpace.Description,
			data.MockSpace.CreatedAt,
			data.MockSpace.UpdatedAt,
		),
	},
	{
		adapter.FindSpacesResponsePayload{
			Spaces: []space.DTO{
				data.MockSpace,
				data.MockSpace,
			},
		},
		fmt.Sprintf(`{"spaces":[{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"},{"uid":"%s","title":"%s","lenderId":"%s","location":{"address":"%s","city":"%s","country":"%s","countryCode":"%s","phone":"%s","province":"%s","provinceCode":"%s","zip":"%s","coordinate":{"latitude":%g,"longitude":%g}},"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s","updatedAt":"%s"}]}`,
			data.MockSpace.UId,
			data.MockSpace.Title,
			data.MockSpace.LenderId,
			data.MockSpace.Location.Address,
			data.MockSpace.Location.City,
			data.MockSpace.Location.Country,
			data.MockSpace.Location.CountryCode,
			data.MockSpace.Location.Phone,
			data.MockSpace.Location.Province,
			data.MockSpace.Location.ProvinceCode,
			data.MockSpace.Location.Zip,
			data.MockSpace.Location.Coordinate.Latitude,
			data.MockSpace.Location.Coordinate.Longitude,
			data.MockSpace.ImageUrls[0],
			data.MockSpace.ImageUrls[1],
			data.MockSpace.Description,
			data.MockSpace.CreatedAt,
			data.MockSpace.UpdatedAt,
			data.MockSpace.UId,
			data.MockSpace.Title,
			data.MockSpace.LenderId,
			data.MockSpace.Location.Address,
			data.MockSpace.Location.City,
			data.MockSpace.Location.Country,
			data.MockSpace.Location.CountryCode,
			data.MockSpace.Location.Phone,
			data.MockSpace.Location.Province,
			data.MockSpace.Location.ProvinceCode,
			data.MockSpace.Location.Zip,
			data.MockSpace.Location.Coordinate.Latitude,
			data.MockSpace.Location.Coordinate.Longitude,
			data.MockSpace.ImageUrls[0],
			data.MockSpace.ImageUrls[1],
			data.MockSpace.Description,
			data.MockSpace.CreatedAt,
			data.MockSpace.UpdatedAt,
		),
	},
	{
		adapter.AddSpaceResponsePayload{
			UId: data.MockSpace.UId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, data.MockSpace.UId),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
