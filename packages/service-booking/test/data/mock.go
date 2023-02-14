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
	MockStartDate     = faker.Date()
	MockEndDate       = faker.Date()
	MockCreatedAt     = faker.Date()
	MockUpdatedAt     = faker.Date()
	MockUserId        = faker.UUIDDigit()
	MockSpaceId       = faker.UUIDDigit()
	MockImageUrls     = []string{faker.URL(), faker.URL()}
	MockDescription   = "BLAH BLAH BLAH"
	MockBookingEntity = booking.DTO{
		UId:         MockUId,
		Status:      "PENDING",
		ImageUrls:   MockImageUrls,
		UserId:      MockUserId,
		SpaceId:     MockSpaceId,
		Description: MockDescription,
		StartDate:   MockStartDate,
		EndDate:     MockEndDate,
		CreatedAt:   MockCreatedAt,
		UpdatedAt:   MockUpdatedAt,
	}.ToEntity()
)
