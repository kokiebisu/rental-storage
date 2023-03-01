// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	mock "github.com/stretchr/testify/mock"
)

// ItemService is an autogenerated mock type for the ItemService type
type ItemService struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: name, imageUrls, ownerId, spaceId
func (_m *ItemService) AddItem(name string, imageUrls []string, ownerId string, spaceId string) *errors.CustomError {
	ret := _m.Called(name, imageUrls, ownerId, spaceId)

	var r0 *errors.CustomError
	if rf, ok := ret.Get(0).(func(string, []string, string, string) *errors.CustomError); ok {
		r0 = rf(name, imageUrls, ownerId, spaceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.CustomError)
		}
	}

	return r0
}

type mockConstructorTestingTNewItemService interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemService creates a new instance of ItemService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemService(t mockConstructorTestingTNewItemService) *ItemService {
	mock := &ItemService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
