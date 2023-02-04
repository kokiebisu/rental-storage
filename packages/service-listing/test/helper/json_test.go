package test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	"github.com/kokiebisu/rental-storage/service-listing/internal/helper"
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
	latitude      = float32(faker.Longitude())
	longitude     = float32(faker.Latitude())
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
	l = listing.DTO{
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
		controller.FindListingByIdResponsePayload{
			Listing: l,
		},
		fmt.Sprintf(`{"listing":{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}}}`,
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
		controller.FindListingsWithinLatLngResponsePayload{
			Listings: []listing.DTO{
				l,
				l,
			},
		},
		fmt.Sprintf(`{"listings":[{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}},{"uid":"%s","title":"%s","lenderId":"%s","streetAddress":"%s","latitude":%g,"longitude":%g,"imageUrls":["%s","%s"],"fee":{"amount":{"value":%d,"currency":"%s"},"type":"%s"}}]}`,
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
		controller.AddListingResponsePayload{
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
