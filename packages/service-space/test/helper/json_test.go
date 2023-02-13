package test

import (
	"fmt"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/controller"
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
		controller.FindSpaceByIdResponsePayload{
			Space: data.MockSpace,
		},
		fmt.Sprintf(`{"space":{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s"}}`,
			data.MockUId,
			data.MockTitle,
			data.MockLenderId,
			data.MockStreetAddress,
			data.MockLatitude,
			data.MockLongitude,
			data.MockImageUrls[0],
			data.MockImageUrls[1],
			data.MockDescription,
			data.MockDate,
		),
	},
	{
		controller.FindSpacesResponsePayload{
			Spaces: []space.DTO{
				data.MockSpace,
				data.MockSpace,
			},
		},
		fmt.Sprintf(`{"spaces":[{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s"},{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"description":"%s","createdAt":"%s"}]}`,
			data.MockUId,
			data.MockTitle,
			data.MockLenderId,
			data.MockStreetAddress,
			data.MockLatitude,
			data.MockLongitude,
			data.MockImageUrls[0],
			data.MockImageUrls[1],
			data.MockDescription,
			data.MockDate,
			data.MockUId,
			data.MockTitle,
			data.MockLenderId,
			data.MockStreetAddress,
			data.MockLatitude,
			data.MockLongitude,
			data.MockImageUrls[0],
			data.MockImageUrls[1],
			data.MockDescription,
			data.MockDate,
		),
	},
	{
		controller.AddSpaceResponsePayload{
			UId: data.MockUId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, data.MockUId),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
