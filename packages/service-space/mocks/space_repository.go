// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	errors "github.com/kokiebisu/rental-storage/service-space/internal/error"
	mock "github.com/stretchr/testify/mock"

	space "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
)

// SpaceRepository is an autogenerated mock type for the SpaceRepository type
type SpaceRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: uid
func (_m *SpaceRepository) Delete(uid string) (string, *errors.CustomError) {
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

// FindManyByLatLng provides a mock function with given fields: latitude, longitude, distance
func (_m *SpaceRepository) FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]space.Entity, *errors.CustomError) {
	ret := _m.Called(latitude, longitude, distance)

	var r0 []space.Entity
	if rf, ok := ret.Get(0).(func(float64, float64, int32) []space.Entity); ok {
		r0 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]space.Entity)
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

// FindManyByUserId provides a mock function with given fields: userId
func (_m *SpaceRepository) FindManyByUserId(userId string) ([]space.Entity, *errors.CustomError) {
	ret := _m.Called(userId)

	var r0 []space.Entity
	if rf, ok := ret.Get(0).(func(string) []space.Entity); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]space.Entity)
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

// FindOneById provides a mock function with given fields: uid
func (_m *SpaceRepository) FindOneById(uid string) (space.Entity, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 space.Entity
	if rf, ok := ret.Get(0).(func(string) space.Entity); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(space.Entity)
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

// Save provides a mock function with given fields: _a0
func (_m *SpaceRepository) Save(_a0 space.Entity) (string, *errors.CustomError) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(space.Entity) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(space.Entity) *errors.CustomError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// Setup provides a mock function with given fields:
func (_m *SpaceRepository) Setup() *errors.CustomError {
	ret := _m.Called()

	var r0 *errors.CustomError
	if rf, ok := ret.Get(0).(func() *errors.CustomError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.CustomError)
		}
	}

	return r0
}

type mockConstructorTestingTNewSpaceRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpaceRepository creates a new instance of SpaceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpaceRepository(t mockConstructorTestingTNewSpaceRepository) *SpaceRepository {
	mock := &SpaceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}