// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	location "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	errors "github.com/kokiebisu/rental-storage/service-space/internal/error"
	mock "github.com/stretchr/testify/mock"

	space "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
)

// SpaceService is an autogenerated mock type for the SpaceService type
type SpaceService struct {
	mock.Mock
}

// CreateSpace provides a mock function with given fields: uid, lenderId, _a2, imageUrls, title, description, createdAt, updatedAt
func (_m *SpaceService) CreateSpace(uid string, lenderId string, _a2 location.DTO, imageUrls []string, title string, description string, createdAt string, updatedAt string) (string, *errors.CustomError) {
	ret := _m.Called(uid, lenderId, _a2, imageUrls, title, description, createdAt, updatedAt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, location.DTO, []string, string, string, string, string) string); ok {
		r0 = rf(uid, lenderId, _a2, imageUrls, title, description, createdAt, updatedAt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string, string, location.DTO, []string, string, string, string, string) *errors.CustomError); ok {
		r1 = rf(uid, lenderId, _a2, imageUrls, title, description, createdAt, updatedAt)
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
