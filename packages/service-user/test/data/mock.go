package data

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

var (
	MockUser = user.DTO{
		UId:          faker.UUIDDigit(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
		EmailAddress: faker.Email(),
		Password:     faker.Password(),
		Items:        []item.DTO{},
		CreatedAt:    "2006-01-02 15:04:05",
		UpdatedAt:    "2006-01-02 15:04:05",
	}
)
