package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
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
	MockDescription = "BLAH BLAH"
	MockSpace       = space.Entity{
		UId:           MockUId,
		Title:         MockTitle,
		LenderId:      MockLenderId,
		StreetAddress: streetaddress.ValueObject{Value: MockStreetAddress},
		Latitude:      coordinate.ValueObject{Value: MockLatitude},
		Longitude:     coordinate.ValueObject{Value: MockLongitude},
		ImageUrls:     MockImageUrls,
		Description:   MockDescription,
	}
)
