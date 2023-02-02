// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	amount "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

	fee "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"

	listing "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"

	mock "github.com/stretchr/testify/mock"
)

// ListingService is an autogenerated mock type for the ListingService type
type ListingService struct {
	mock.Mock
}

// CreateListing provides a mock function with given fields: lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType
func (_m *ListingService) CreateListing(lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *errors.CustomError) {
	ret := _m.Called(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, float32, float32, []string, string, int32, amount.CurrencyType, fee.RentalFeeType) string); ok {
		r0 = rf(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string, string, float32, float32, []string, string, int32, amount.CurrencyType, fee.RentalFeeType) *errors.CustomError); ok {
		r1 = rf(lenderId, streetAddress, latitude, longitude, imageUrls, title, feeAmount, feeCurrency, feeType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// FindListingById provides a mock function with given fields: uid
func (_m *ListingService) FindListingById(uid string) (listing.DTO, *errors.CustomError) {
	ret := _m.Called(uid)

	var r0 listing.DTO
	if rf, ok := ret.Get(0).(func(string) listing.DTO); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(listing.DTO)
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

// FindListingsWithinLatLng provides a mock function with given fields: latitude, longitude, distance
func (_m *ListingService) FindListingsWithinLatLng(latitude float32, longitude float32, distance int32) ([]listing.DTO, *errors.CustomError) {
	ret := _m.Called(latitude, longitude, distance)

	var r0 []listing.DTO
	if rf, ok := ret.Get(0).(func(float32, float32, int32) []listing.DTO); ok {
		r0 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]listing.DTO)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(float32, float32, int32) *errors.CustomError); ok {
		r1 = rf(latitude, longitude, distance)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// RemoveListingById provides a mock function with given fields: uid
func (_m *ListingService) RemoveListingById(uid string) *errors.CustomError {
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

type mockConstructorTestingTNewListingService interface {
	mock.TestingT
	Cleanup(func())
}

// NewListingService creates a new instance of ListingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListingService(t mockConstructorTestingTNewListingService) *ListingService {
	mock := &ListingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}