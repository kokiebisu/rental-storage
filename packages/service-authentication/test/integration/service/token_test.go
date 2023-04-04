package integration

import (
	"testing"
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken_Success(t *testing.T) {
	token, err := data.TokenService.GenerateAccessToken(data.MockUId, time.Hour*1)
	assert.Nil(t, err, "should not throw error")
	assert.NotEmpty(t, token, "should not be empty")
}

func TestGenerateRefreshToken_Success(t *testing.T) {
	token, err := data.TokenService.GenerateRefreshToken(data.MockUId, time.Hour*1)
	assert.Nil(t, err, "should not throw error")
	assert.NotEmpty(t, token, "should not be empty")
}

func TestVerifyToken_Success(t *testing.T) {
	token, err := data.TokenService.GenerateAccessToken(data.MockUId, time.Hour*1)
	assert.Nil(t, err, "should not throw error")
	claims, err := data.TokenService.VerifyToken(string(token))
	assert.Nil(t, err, "should not throw error")
	assert.NotNil(t, claims, "should not be nil")
}
