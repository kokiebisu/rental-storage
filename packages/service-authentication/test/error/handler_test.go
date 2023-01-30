package test

import (
	"errors"
	"testing"

	errs "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/stretchr/testify/assert"
)

// for a valid return value.
func TestCustomError(t *testing.T) {
	msg := "something went wrong"
	err := errors.New(msg)
	customError := errs.ErrorHandler.CustomError(msg, err)
	assert.Equal(t, uint16(500), customError.StatusCode)
	assert.Equal(t, msg, customError.Error())
}
