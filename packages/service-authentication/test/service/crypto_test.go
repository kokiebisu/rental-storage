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

}
