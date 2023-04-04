// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// GenerateAccessToken provides a mock function with given fields: uid, expiresAt
func (_m *TokenService) GenerateAccessToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError) {
	ret := _m.Called(uid, expiresAt)

	var r0 domain.Token
	var r1 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(string, time.Duration) (domain.Token, *customerror.CustomError)); ok {
		return rf(uid, expiresAt)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration) domain.Token); ok {
		r0 = rf(uid, expiresAt)
	} else {
		r0 = ret.Get(0).(domain.Token)
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration) *customerror.CustomError); ok {
		r1 = rf(uid, expiresAt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

// GenerateRefreshToken provides a mock function with given fields: uid, expiresAt
func (_m *TokenService) GenerateRefreshToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError) {
	ret := _m.Called(uid, expiresAt)

	var r0 domain.Token
	var r1 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(string, time.Duration) (domain.Token, *customerror.CustomError)); ok {
		return rf(uid, expiresAt)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration) domain.Token); ok {
		r0 = rf(uid, expiresAt)
	} else {
		r0 = ret.Get(0).(domain.Token)
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration) *customerror.CustomError); ok {
		r1 = rf(uid, expiresAt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: tokenString
func (_m *TokenService) VerifyToken(tokenString string) (*domain.Claims, *customerror.CustomError) {
	ret := _m.Called(tokenString)

	var r0 *domain.Claims
	var r1 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(string) (*domain.Claims, *customerror.CustomError)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Claims); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Claims)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *customerror.CustomError); ok {
		r1 = rf(tokenString)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewTokenService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenService creates a new instance of TokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenService(t mockConstructorTestingTNewTokenService) *TokenService {
	mock := &TokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
