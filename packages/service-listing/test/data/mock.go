package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/coordinate"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/street_address"
)

var (
	MockUId           = faker.UUIDDigit()
	MockTitle         = "A beautiful garage"
	MockLenderId      = faker.UUIDDigit()
	MockStreetAddress = faker.FirstName()
	MockLatitude      = float32(faker.Latitude())
	MockLongitude     = float32(faker.Longitude())
	MockDistance      = int32(5)
	MockImageUrls     = []string{
		faker.URL(),
		faker.URL(),
	}
	MockFeeCurrency = amount.CurrencyType("CAD")
	MockFeeAmount   = int64(50)
	MockFeeType     = "MONTHLY"
	MockListing     = listing.Entity{
		UId:           MockUId,
		Title:         MockTitle,
		LenderId:      MockLenderId,
		StreetAddress: streetaddress.ValueObject{Value: MockStreetAddress},
		Latitude:      coordinate.ValueObject{Value: MockLatitude},
		Longitude:     coordinate.ValueObject{Value: MockLongitude},
		ImageUrls:     MockImageUrls,
		Fee: fee.ValueObject{
			Amount: amount.ValueObject{
				Value:    MockFeeAmount,
				Currency: MockFeeCurrency,
			},
			Type: "MONTHLY",
		},
	}
)
