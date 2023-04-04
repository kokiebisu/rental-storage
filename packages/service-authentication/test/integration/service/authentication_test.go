package integration

import (
	"testing"
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/stretchr/testify/assert"
)

func TestSignIn_Success(t *testing.T) {
	tokens, err := data.AuthenticationService.SignIn(data.MockEmailAddress, data.MockPassword)
	assert.Nil(t, err, "should not throw error")

	assert.NotEmpty(t, tokens["access_token"], "should not be empty")
	assert.NotEmpty(t, tokens["refresh_token"], "should not be empty")
}

func TestSignUp_Success(t *testing.T) {
	token, err := data.AuthenticationService.SignUp(data.MockEmailAddress, data.MockFirstName, data.MockLastName, data.MockPassword)
	assert.Nil(t, err, "should not throw error")

	assert.NotEmpty(t, token, "should not be empty")
}

func TestVerify_Success(t *testing.T) {
	token, err := data.TokenService.GenerateAccessToken(data.MockUId, time.Hour*1)
	assert.Nil(t, err, "should not throw error")
	claims, err := data.AuthenticationService.Verify(string(token))
	assert.Nil(t, err, "should not throw error")
	assert.NotNil(t, claims, "should not be nil")
}
