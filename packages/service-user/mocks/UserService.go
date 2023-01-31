// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	item "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"

	mock "github.com/stretchr/testify/mock"

	user "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: uid, emailAddress, firstName, lastName, password, items, createdAt
func (_m *UserService) CreateUser(uid string, emailAddress string, firstName string, lastName string, password string, items []item.DTO, createdAt string) (string, *errors.CustomError) {
	ret := _m.Called(uid, emailAddress, firstName, lastName, password, items, createdAt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, []item.DTO, string) string); ok {
		r0 = rf(uid, emailAddress, firstName, lastName, password, items, createdAt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string, string, string, string, string, []item.DTO, string) *errors.CustomError); ok {
		r1 = rf(uid, emailAddress, firstName, lastName, password, items, createdAt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: emailAddress
func (_m *UserService) FindByEmail(emailAddress string) (user.Entity, *errors.CustomError) {
	ret := _m.Called(emailAddress)

	var r0 user.Entity
	if rf, ok := ret.Get(0).(func(string) user.Entity); ok {
		r0 = rf(emailAddress)
	} else {
		r0 = ret.Get(0).(user.Entity)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(emailAddress)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindById provides a mock function with given fields: uid
func (_m *UserService) FindById(uid string) (user.Entity, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 user.Entity
	if rf, ok := ret.Get(0).(func(string) user.Entity); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(user.Entity)
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

// RemoveById provides a mock function with given fields: uid
func (_m *UserService) RemoveById(uid string) *errors.CustomError {
	ret := _m.Called(uid)

	var r0 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) *errors.CustomError); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.CustomError)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}