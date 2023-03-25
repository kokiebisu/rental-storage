package data

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang-jwt/jwt"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
)

var (
	MockUId            = faker.UUIDDigit()
	MockEmailAddress   = faker.Email()
	MockFirstName      = faker.Name()
	MockLastName       = faker.Name()
	MockPassword       = "1234"
	MockHashedPassword = "$2a$10$6Ci1sE7FWKCU5jZAumHQ2e5306sKwDkrnvHqAozIRfEMBFNjPq4Ce" // hashed from '1234'
	MockToken          = faker.Jwt()
	MockTimeString     = faker.TimeString()
	MockExpiresAt      = time.Duration(5 * time.Minute)
	MockClaims         = &domain.Claims{
		MockUId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}
)
