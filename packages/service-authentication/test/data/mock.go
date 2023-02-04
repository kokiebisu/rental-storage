package data

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang-jwt/jwt"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
)

var (
	MockUId          = faker.UUIDDigit()
	MockEmailAddress = faker.Email()
	MockFirstName    = faker.Name()
	MockLastName     = faker.Name()
	MockPassword     = faker.Password()
	MockToken        = faker.Jwt()
	MockTimeString   = faker.TimeString()
	MockClaims       = &domain.Claims{
		MockUId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}
)
