package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken_Success(t *testing.T) {
	token, err := data.TokenService.GenerateToken(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	assert.NotEmpty(t, token, "should not be empty")
}

func TestVerifyToken_Success(t *testing.T) {
	claims, err := data.TokenService.VerifyToken(data.MockToken)
	assert.Nil(t, err, "should not throw error")
	assert.NotNil(t, claims, "should not be nil")
}
