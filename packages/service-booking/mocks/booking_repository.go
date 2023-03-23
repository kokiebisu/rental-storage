// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	booking "github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"

	mock "github.com/stretchr/testify/mock"
)

// BookingRepository is an autogenerated mock type for the BookingRepository type
type BookingRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *BookingRepository) Delete(id string) *errors.CustomError {
	ret := _m.Called(id)

	var r0 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) *errors.CustomError); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.CustomError)
		}
	}

	return r0
}

// FindManyBySpaceId provides a mock function with given fields: spaceId, status
func (_m *BookingRepository) FindManyBySpaceId(spaceId string, status string) ([]booking.Entity, *errors.CustomError) {
	ret := _m.Called(spaceId, status)

	var r0 []booking.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string, string) ([]booking.Entity, *errors.CustomError)); ok {
		return rf(spaceId, status)
	}
	if rf, ok := ret.Get(0).(func(string, string) []booking.Entity); ok {
		r0 = rf(spaceId, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Entity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *errors.CustomError); ok {
		r1 = rf(spaceId, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindManyByUserId provides a mock function with given fields: userId, status
func (_m *BookingRepository) FindManyByUserId(userId string, status string) ([]booking.Entity, *errors.CustomError) {
	ret := _m.Called(userId, status)

	var r0 []booking.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string, string) ([]booking.Entity, *errors.CustomError)); ok {
		return rf(userId, status)
	}
	if rf, ok := ret.Get(0).(func(string, string) []booking.Entity); ok {
		r0 = rf(userId, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Entity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *errors.CustomError); ok {
		r1 = rf(userId, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindOneById provides a mock function with given fields: id
func (_m *BookingRepository) FindOneById(id string) (booking.Entity, *errors.CustomError) {
	ret := _m.Called(id)

	var r0 booking.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) (booking.Entity, *errors.CustomError)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) booking.Entity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(booking.Entity)
	}

	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *BookingRepository) Save(_a0 booking.Entity) *errors.CustomError {
	ret := _m.Called(_a0)

	var r0 *errors.CustomError
	if rf, ok := ret.Get(0).(func(booking.Entity) *errors.CustomError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.CustomError)
		}
	}

	return r0
}

// UpdateBookingStatus provides a mock function with given fields: id, status
func (_m *BookingRepository) UpdateBookingStatus(id string, status string) (booking.Entity, *errors.CustomError) {
	ret := _m.Called(id, status)

	var r0 booking.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string, string) (booking.Entity, *errors.CustomError)); ok {
		return rf(id, status)
	}
	if rf, ok := ret.Get(0).(func(string, string) booking.Entity); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Get(0).(booking.Entity)
	}

	if rf, ok := ret.Get(1).(func(string, string) *errors.CustomError); ok {
		r1 = rf(id, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewBookingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingRepository creates a new instance of BookingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingRepository(t mockConstructorTestingTNewBookingRepository) *BookingRepository {
	mock := &BookingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
