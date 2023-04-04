package unit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

// Generate Token
func TestGenerateAccessToken(t *testing.T) {
	tokenService := service.NewTokenService()
	token, err := tokenService.GenerateAccessToken(data.MockUId, time.Hour*1)
	if err != nil {
		panic("something went wrong")
	}
	assert.Greater(t, len(token), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

func TestGenerateRefreshToken(t *testing.T) {
	tokenService := service.NewTokenService()
	token, err := tokenService.GenerateRefreshToken(data.MockUId, time.Hour*1)
	if err != nil {
		panic("something went wrong")
	}
	assert.Greater(t, len(token), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

// VerifyToken
func TestVerifyToken_Success(t *testing.T) {
	tokenService := service.NewTokenService()
	token, err := tokenService.GenerateAccessToken(data.MockUId, time.Hour*1)
	assert.Nil(t, err, "should be no errors")
	claims, err := tokenService.VerifyToken(string(token))
	assert.Greater(t, len(claims.UId), 0, "should be valid uid")
	assert.Nil(t, err, "should be no errors")
}
