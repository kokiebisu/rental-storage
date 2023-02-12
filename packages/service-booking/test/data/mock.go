package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
)

var (
	MockUId          = faker.UUIDDigit()
	MockEmailAddress = faker.Email()
	MockFirstName    = faker.FirstName()
	MockLastName     = faker.LastName()
	MockPassword     = faker.Password()
	MockDateString   = faker.Date()
	MockAmount       = amount.DTO{
		Value:    50,
		Currency: faker.Currency(),
	}
	MockUserId  = faker.UUIDDigit()
	MockSpaceId = faker.UUIDDigit()
	MockItem    = item.DTO{
		Id:        111111,
		Name:      faker.Name(),
		ImageUrls: []string{faker.URL(), faker.URL()},
	}
	MockBookingEntity = booking.DTO{
		UId:       MockUId,
		Status:    "PENDING",
		UserId:    MockUserId,
		SpaceId:   MockSpaceId,
		Items:     []item.DTO{MockItem},
		CreatedAt: MockDateString,
		UpdatedAt: MockDateString,
	}.ToEntity()
)
