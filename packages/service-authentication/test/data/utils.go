package data

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient           *redis.Client
	TokenStore            port.TokenStore
	AuthenticationService port.AuthenticationService
	TokenService          port.TokenService
	CryptoService         port.CryptoService
)
