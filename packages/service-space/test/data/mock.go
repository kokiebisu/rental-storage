package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
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
	MockSpace       = space.DTO{
		UId:           MockUId,
		Title:         MockTitle,
		LenderId:      MockLenderId,
		StreetAddress: MockStreetAddress,
		Latitude:      MockLatitude,
		Longitude:     MockLongitude,
		ImageUrls:     MockImageUrls,
		Description:   MockDescription,
		CreatedAt:     MockDate,
	}
	MockDate = faker.Date()
)
