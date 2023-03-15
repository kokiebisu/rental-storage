package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
)

var (
	MockEmailAddress = faker.Email()
	MockFirstName    = faker.FirstName()
	MockLastName     = faker.LastName()
	MockPassword     = faker.Password()
	MockBooking      = booking.DTO{
		UId:           faker.UUIDDigit(),
		BookingStatus: "pending",
		ImageUrls:     []string{faker.URL(), faker.URL()},
		UserId:        faker.UUIDDigit(),
		SpaceId:       faker.UUIDDigit(),
		Description:   "BLAH BLAH BLAH",
		StartDate:     faker.Date(),
		EndDate:       faker.Date(),
		CreatedAt:     faker.Date(),
		UpdatedAt:     faker.Date(),
	}
)
