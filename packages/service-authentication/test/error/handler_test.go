package test

import (
	"errors"
	"testing"

	errs "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalError(t *testing.T) {
	target := "body"
	msg := "unable to unmarshal " + target
	err := errors.New(msg)
	customError := errs.ErrorHandler.UnmarshalError(target, err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestRequestFailError(t *testing.T){
	msg := "failed to send request to user service endpoint"
	err := errors.New(msg)
	customError := errs.ErrorHandler.RequestFailError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestDecodeError(t *testing.T) {
	target := "user service endpoint to user domain"
	msg := "unable to decode " + target
	err := errors.New(msg)
	customError := errs.ErrorHandler.DecodeError("user service endpoint to user domain", err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestPasswordHashErrorError(t *testing.T){
	msg := "cannot hash password"
	err := errors.New(msg)
	customError := errs.ErrorHandler.PasswordHashError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestUndefinedEndPointError(t *testing.T){
	msg := "user service api endpoint not defined"
	err := errors.New(msg)
	customError := errs.ErrorHandler.UndefinedEndPointError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestResponseInvalidError(t *testing.T){
	msg := "response from user service was invalid"
	err := errors.New(msg)
	customError := errs.ErrorHandler.ResponseInvalidError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestSignTokeError(t *testing.T){
	msg := "unable to sign token"
	err := errors.New(msg)
	customError := errs.ErrorHandler.SignTokenError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestClaimParseError(t *testing.T){
	msg := "unable to parse with claims"
	err := errors.New(msg)
	customError := errs.ErrorHandler.ClaimParseError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestClaimCastError(t *testing.T){
	msg := "unable to cast to domain claims"
	err := errors.New(msg)
	customError := errs.ErrorHandler.ClaimCastError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestClaimUidEmptyError(t *testing.T){
	msg := "uid property in claims is empty string"
	err := errors.New(msg)
	customError := errs.ErrorHandler.ClaimUidEmptyError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestClaimExpiredError(t *testing.T){
	msg := "claims should have been expired already"
	err := errors.New(msg)
	customError := errs.ErrorHandler.ClaimExpiredError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}

func TestCompareHashError(t *testing.T){
	msg := "failed when comparing hash and password"
	err := errors.New(msg)
	customError := errs.ErrorHandler.CompareHashError(err)
	payload := customError.GetPayload()
	assert.Equal(t, customError.StatusCode, uint16(500))
	assert.Equal(t, payload.Message, msg)
	assert.Equal(t, payload.Reason, err)
}
