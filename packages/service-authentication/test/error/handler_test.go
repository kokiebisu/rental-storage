package test

import (
	"testing"

	errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/stretchr/testify/assert"
)

// for a valid return value.
func TestCustomError(t *testing.T) {
	msg := "something went wrong"
	err := errors.ErrorHandler.CustomError(msg)
	assert.Equal(t, uint16(500), err.StatusCode)
	assert.Equal(t, msg, err.Error())
}
