package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
)

var (
	MockUId           = faker.UUIDDigit()
	MockEmailAddress  = faker.Email()
	MockFirstName     = faker.FirstName()
	MockLastName      = faker.LastName()
	MockPassword      = faker.Password()
	MockDateString    = faker.Date()
	MockUserId        = faker.UUIDDigit()
	MockSpaceId       = faker.UUIDDigit()
	MockBookingEntity = booking.DTO{
		UId:       MockUId,
		Status:    "PENDING",
		UserId:    MockUserId,
		SpaceId:   MockSpaceId,
		CreatedAt: MockDateString,
		UpdatedAt: MockDateString,
	}.ToEntity()
)
