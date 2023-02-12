// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	amount "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	errors "github.com/kokiebisu/rental-storage/service-space/internal/error"

	mock "github.com/stretchr/testify/mock"

	space "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
)

// SpaceFactory is an autogenerated mock type for the SpaceFactory type
type SpaceFactory struct {
	mock.Mock
}

// New provides a mock function with given fields: title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, feeAmount, feeType
func (_m *SpaceFactory) New(title string, lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, feeCurrency amount.CurrencyType, feeAmount int64, feeType string) (space.Entity, *errors.CustomError) {
	ret := _m.Called(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, feeAmount, feeType)

	var r0 space.Entity
	if rf, ok := ret.Get(0).(func(string, string, string, float64, float64, []string, amount.CurrencyType, int64, string) space.Entity); ok {
		r0 = rf(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, feeAmount, feeType)
	} else {
		r0 = ret.Get(0).(space.Entity)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string, string, string, float64, float64, []string, amount.CurrencyType, int64, string) *errors.CustomError); ok {
		r1 = rf(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, feeAmount, feeType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewSpaceFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpaceFactory creates a new instance of SpaceFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpaceFactory(t mockConstructorTestingTNewSpaceFactory) *SpaceFactory {
	mock := &SpaceFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}