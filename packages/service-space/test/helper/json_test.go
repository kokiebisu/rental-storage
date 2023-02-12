package test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"
	"github.com/kokiebisu/rental-storage/service-space/internal/helper"
	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var (
	uid           = faker.UUIDDigit()
	title         = faker.Name()
	lenderId      = faker.FirstName()
	streetAddress = faker.Word()
	latitude      = faker.Longitude()
	longitude     = faker.Latitude()
	imageUrls     = []string{
		faker.URL(),
		faker.URL(),
	}
	f = fee.DTO{
		Amount: amount.DTO{
			Value:    50,
			Currency: faker.Currency(),
		},
		Type: "MONTHLY",
	}
	l = space.DTO{
		UId:           uid,
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: streetAddress,
		Latitude:      latitude,
		Longitude:     longitude,
		ImageUrls:     imageUrls,
		Fee:           f,
	}
)

var stringifyResults = []StringifyResult{
	{
		controller.FindSpaceByIdResponsePayload{
			Space: l,
		},
		fmt.Sprintf(`{"space":{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}}}`,
			uid,
			title,
			lenderId,
			streetAddress,
			latitude,
			longitude,
			imageUrls[0],
			imageUrls[1],
			f.Amount.Value,
			f.Amount.Currency,
			f.Type,
		),
	},
	{
		controller.FindSpacesResponsePayload{
			Spaces: []space.DTO{
				l,
				l,
			},
		},
		fmt.Sprintf(`{"spaces":[{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}},{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}}]}`,
			uid,
			title,
			lenderId,
			streetAddress,
			latitude,
			longitude,
			imageUrls[0],
			imageUrls[1],
			f.Amount.Value,
			f.Amount.Currency,
			f.Type,
			uid,
			title,
			lenderId,
			streetAddress,
			latitude,
			longitude,
			imageUrls[0],
			imageUrls[1],
			f.Amount.Value,
			f.Amount.Currency,
			f.Type,
		),
	},
	{
		controller.AddSpaceResponsePayload{
			UId: uid,
		},
		fmt.Sprintf(`{"uid":"%s"}`, uid),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}