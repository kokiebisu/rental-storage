// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	amount "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	errors "github.com/kokiebisu/rental-storage/service-space/internal/error"

	fee "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"

	mock "github.com/stretchr/testify/mock"

	space "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
)

// SpaceService is an autogenerated mock type for the SpaceService type
type SpaceService struct {
	mock.Mock
}

// CreateSpace provides a mock function with given fields: lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType
func (_m *SpaceService) CreateSpace(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *errors.CustomError) {
	ret := _m.Called(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, float64, float64, []string, string, int32, amount.CurrencyType, fee.RentalFeeType) string); ok {
		r0 = rf(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string, string, float64, float64, []string, string, int32, amount.CurrencyType, fee.RentalFeeType) *errors.CustomError); ok {
		r1 = rf(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindSpaceById provides a mock function with given fields: uid
func (_m *SpaceService) FindSpaceById(uid string) (space.DTO, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 space.DTO
	if rf, ok := ret.Get(0).(func(string) space.DTO); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(space.DTO)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(uid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindSpacesByUserId provides a mock function with given fields: userId
func (_m *SpaceService) FindSpacesByUserId(userId string) ([]space.DTO, *errors.CustomError) {
	ret := _m.Called(userId)

	var r0 []space.DTO
	if rf, ok := ret.Get(0).(func(string) []space.DTO); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]space.DTO)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindSpacesWithinLatLng provides a mock function with given fields: latitude, longitude, distance
func (_m *SpaceService) FindSpacesWithinLatLng(latitude float64, longitude float64, distance int32) ([]space.DTO, *errors.CustomError) {
	ret := _m.Called(latitude, longitude, distance)

	var r0 []space.DTO
	if rf, ok := ret.Get(0).(func(float64, float64, int32) []space.DTO); ok {
		r0 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]space.DTO)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(float64, float64, int32) *errors.CustomError); ok {
		r1 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// RemoveSpaceById provides a mock function with given fields: uid
func (_m *SpaceService) RemoveSpaceById(uid string) (string, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(uid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewSpaceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpaceService creates a new instance of SpaceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpaceService(t mockConstructorTestingTNewSpaceService) *SpaceService {
	mock := &SpaceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}