package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/stretchr/testify/assert"
)

func TestAccessToken_Success(t *testing.T) {
	err := data.TokenStore.SetAccessToken(data.MockUId, data.MockToken, data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Get access token
	retrievedToken, err := data.TokenStore.GetAccessToken(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, data.MockToken, retrievedToken, "should be equal")
}

func TestRefreshToken_Success(t *testing.T) {
	err := data.TokenStore.SetRefreshToken(data.MockUId, data.MockToken, data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Get refresh token
	token, err := data.TokenStore.GetRefreshToken(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, data.MockToken, token, "should be equal")
}

func TestDeleteAccessToken_Success(t *testing.T) {
	err := data.TokenStore.SetAccessToken(data.MockUId, data.MockToken, data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Delete access token
	err = data.TokenStore.DeleteAccessToken(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	// Get access token
	token, err := data.TokenStore.GetAccessToken(data.MockUId)
	assert.NotNil(t, err, "should throw error")
	assert.Equal(t, "", token, "should be empty")
}

func TestDeleteRefreshToken_Success(t *testing.T) {
	err := data.TokenStore.SetRefreshToken(data.MockUId, data.MockToken, data.MockExpiresAt)
	assert.Nil(t, err, "should not throw error")
	// Delete refresh token
	err = data.TokenStore.DeleteRefreshToken(data.MockUId)
	assert.Nil(t, err, "should not throw error")
	// Get refresh token
	token, err := data.TokenStore.GetRefreshToken(data.MockUId)
	assert.NotNil(t, err, "should throw error")
	assert.Equal(t, "", token, "should be empty")
}
