// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: event
func (_m *Controller) CreateUser(event interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(event)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(event)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindUserByEmail provides a mock function with given fields: event
func (_m *Controller) FindUserByEmail(event interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(event)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(event)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindUserById provides a mock function with given fields: event
func (_m *Controller) FindUserById(event interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(event)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(event)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// RemoveUserById provides a mock function with given fields: event
func (_m *Controller) RemoveUserById(event interface{}) (interface{}, *errors.CustomError) {
	ret := _m.Called(event)

	var r0 interface{}
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, *errors.CustomError)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) *errors.CustomError); ok {
		r1 = rf(event)
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
