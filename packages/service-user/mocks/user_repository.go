// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	mock "github.com/stretchr/testify/mock"

	user "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: uid
func (_m *UserRepository) Delete(uid string) (string, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 string
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) (string, *errors.CustomError)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(uid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: emailAddress
func (_m *UserRepository) FindByEmail(emailAddress string) (user.Entity, *errors.CustomError) {
	ret := _m.Called(emailAddress)

	var r0 user.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) (user.Entity, *errors.CustomError)); ok {
		return rf(emailAddress)
	}
	if rf, ok := ret.Get(0).(func(string) user.Entity); ok {
		r0 = rf(emailAddress)
	} else {
		r0 = ret.Get(0).(user.Entity)
	}

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
func (_m *UserRepository) FindById(uid string) (user.Entity, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 user.Entity
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string) (user.Entity, *errors.CustomError)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) user.Entity); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(user.Entity)
	}

	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(uid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// Save provides a mock function with given fields: u
func (_m *UserRepository) Save(u user.Entity) (string, *errors.CustomError) {
	ret := _m.Called(u)

	var r0 string
	var r1 *errors.CustomError
	if rf, ok := ret.Get(0).(func(user.Entity) (string, *errors.CustomError)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(user.Entity) string); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(user.Entity) *errors.CustomError); ok {
		r1 = rf(u)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// Setup provides a mock function with given fields:
func (_m *UserRepository) Setup() *errors.CustomError {
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

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
