package data

import "github.com/bxcodec/faker/v3"

var (
	MockToken     = faker.Jwt()
	MockUId       = faker.UUIDDigit()
	MockExpiresAt = 1000000
)
