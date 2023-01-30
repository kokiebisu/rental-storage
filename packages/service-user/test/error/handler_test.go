package test

import (
	err "errors"
	"testing"

	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	msg := "something went wrong"
	err := errors.ErrorHandler.CustomError(msg, err.New("mocked error"))
	assert.Equal(t, uint16(500), err.StatusCode)
	assert.Equal(t, "mocked error", err.Error())
}
