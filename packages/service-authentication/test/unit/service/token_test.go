package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

// Generate Token
func TestGenerateToken_Success(t *testing.T) {
	tokenService := service.NewTokenService()
	token, err := tokenService.GenerateToken(data.MockUId)
	if err != nil {
		panic("something went wrong")
	}
	assert.Greater(t, len(token), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

// VerifyToken
func TestVerifyToken_Success(t *testing.T) {
	tokenService := service.NewTokenService()
	token, err := tokenService.GenerateToken(data.MockUId)
	assert.Nil(t, err, "should be no errors")
	claims, err := tokenService.VerifyToken(token)
	assert.Greater(t, len(claims.UId), 0, "should be valid uid")
	assert.Nil(t, err, "should be no errors")
}
