package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
)

func TestHashPassword_Success(t *testing.T) {
	cryptoService := service.NewCryptoService()
	password := "random"
	token, err := cryptoService.HashPassword(password)
	assert.Nil(t, err, "should be no errors")
	assert.Greater(t, len(token), 0, "token should have a length greater than 0")
}

func TestVerifyPassword_Success(t *testing.T) {
	cryptoService := service.NewCryptoService()
	password := "random"
	hashed, err := cryptoService.HashPassword(password)
	assert.Nil(t, err, "should be no errors")
	result, err := cryptoService.VerifyPassword(hashed, password)
	assert.Nil(t, err, "should be no errors")
	if err != nil {
		panic("something went wrong")
	}
	assert.True(t, result, "token should have a length greater than 0")
}

func TestVerifyPassword_InvalidPassword(t *testing.T) {
	cryptoService := service.NewCryptoService()
	password := "password"
	hashed, err := cryptoService.HashPassword(password)
	assert.Nil(t, err, "should be no error")
	result, err := cryptoService.VerifyPassword(hashed, "invalid")
	assert.NotNil(t, err, "should be error")
	assert.False(t, result, "result of verify password should be false")
	assert.Equal(t, err.StatusCode, uint16(500))
	assert.Equal(t, err.ErrorCode, "COMPARE_HASH_ERROR")
}
