// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// CreateBooking provides a mock function with given fields: req
func (_m *Controller) CreateBooking(req interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(req)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindBookingById provides a mock function with given fields: req
func (_m *Controller) FindBookingById(req interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(req)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindBookings provides a mock function with given fields: req
func (_m *Controller) FindBookings(req interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(req)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewController interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t mockConstructorTestingTNewController) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}