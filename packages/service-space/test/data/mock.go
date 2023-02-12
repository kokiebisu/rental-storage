package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"
	streetaddress "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/street_address"
)

var (
	MockUId           = faker.UUIDDigit()
	MockTitle         = "A beautiful garage"
	MockLenderId      = faker.UUIDDigit()
	MockStreetAddress = faker.FirstName()
	MockLatitude      = faker.Latitude()
	MockLongitude     = faker.Longitude()
	MockDistance      = int32(5)
	MockImageUrls     = []string{
		faker.URL(),
		faker.URL(),
	}
	MockFeeCurrency = amount.CurrencyType("CAD")
	MockFeeAmount   = int64(50)
	MockFeeType     = "MONTHLY"
	MockSpace       = space.Entity{
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
