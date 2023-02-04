package data

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	emailaddress "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/email_address"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/name"
)

var (
	MockUId          = faker.UUIDDigit()
	MockEmailAddress = faker.Email()
	MockFirstName    = faker.FirstName()
	MockLastName     = faker.LastName()
	MockPassword     = faker.Password()
	MockTimeString   = faker.TimeString()
	MockUser         = user.Entity{
		UId:          MockUId,
		FirstName:    name.ValueObject{Value: MockFirstName, NameType: "firstName"},
		LastName:     name.ValueObject{Value: MockLastName, NameType: "lastName"},
		EmailAddress: emailaddress.ValueObject{Value: MockEmailAddress},
		Password:     MockPassword,
		Items:        []item.Entity{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
)
