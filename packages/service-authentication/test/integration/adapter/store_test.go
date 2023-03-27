package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/stretchr/testify/assert"
)

func TestGetTokens_Success(t *testing.T) {
	err := data.TokenStore.SetAccessToken(data.MockUId, string(data.MockToken), data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	err = data.TokenStore.SetRefreshToken(data.MockUId, string(data.MockToken), data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Get access token
	tokens, err := data.TokenStore.GetTokens(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, string(data.MockToken), tokens["access_token"], "should be equal")
	assert.Equal(t, string(data.MockToken), tokens["refresh_token"], "should be equal")
}

func TestDeleteTokens_Success(t *testing.T) {
	err := data.TokenStore.SetAccessToken(data.MockUId, string(data.MockToken), data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")

	err = data.TokenStore.SetRefreshToken(data.MockUId, string(data.MockToken), data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Delete access token
	err = data.TokenStore.DeleteTokens(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	// Get access token
	tokens, err := data.TokenStore.GetTokens(data.MockUId)
	assert.NotNil(t, err, "should throw error")
	assert.Equal(t, "", tokens["access_token"], "should be empty")
	assert.Equal(t, "", tokens["refresh_token"], "should be empty")
}
