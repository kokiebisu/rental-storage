// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	item "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	mock "github.com/stretchr/testify/mock"

	user "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// NewItem provides a mock function with given fields: name, ownerId, listingId
func (_m *Factory) NewItem(name string, ownerId string, listingId string) item.Entity {
	ret := _m.Called(name, ownerId, listingId)

	var r0 item.Entity
	if rf, ok := ret.Get(0).(func(string, string, string) item.Entity); ok {
		r0 = rf(name, ownerId, listingId)
	} else {
		r0 = ret.Get(0).(item.Entity)
	}

	return r0
}

// NewUser provides a mock function with given fields: uid, firstName, lastName, emailAddress, password, items, createdAt
func (_m *Factory) NewUser(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) user.Entity {
	ret := _m.Called(uid, firstName, lastName, emailAddress, password, items, createdAt)

	var r0 user.Entity
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, []item.Entity, string) user.Entity); ok {
		r0 = rf(uid, firstName, lastName, emailAddress, password, items, createdAt)
	} else {
		r0 = ret.Get(0).(user.Entity)
	}

	return r0
}

type mockConstructorTestingTNewFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFactory(t mockConstructorTestingTNewFactory) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}