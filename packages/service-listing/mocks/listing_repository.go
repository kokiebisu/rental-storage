// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	listing "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

	mock "github.com/stretchr/testify/mock"
)

// ListingRepository is an autogenerated mock type for the ListingRepository type
type ListingRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: uid
func (_m *ListingRepository) Delete(uid string) (string, *errors.CustomError) {
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
func (_m *ListingRepository) FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]listing.Entity, *errors.CustomError) {
	ret := _m.Called(latitude, longitude, distance)

	var r0 []listing.Entity
	if rf, ok := ret.Get(0).(func(float64, float64, int32) []listing.Entity); ok {
		r0 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]listing.Entity)
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
func (_m *ListingRepository) FindManyByUserId(userId string) ([]listing.Entity, *errors.CustomError) {
	ret := _m.Called(userId)

	var r0 []listing.Entity
	if rf, ok := ret.Get(0).(func(string) []listing.Entity); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]listing.Entity)
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
func (_m *ListingRepository) FindOneById(uid string) (listing.Entity, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 listing.Entity
	if rf, ok := ret.Get(0).(func(string) listing.Entity); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(listing.Entity)
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
func (_m *ListingRepository) Save(_a0 listing.Entity) (string, *errors.CustomError) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(listing.Entity) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(listing.Entity) *errors.CustomError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// Setup provides a mock function with given fields:
func (_m *ListingRepository) Setup() *errors.CustomError {
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

type mockConstructorTestingTNewListingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewListingRepository creates a new instance of ListingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListingRepository(t mockConstructorTestingTNewListingRepository) *ListingRepository {
	mock := &ListingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
