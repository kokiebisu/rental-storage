package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/mocks"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
)

// CreateSpace
func TestCreateSpace_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockFactory := mocks.NewSpaceFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockDescription).Return(data.MockSpace, nil)
	mockRepo.On("Save", data.MockSpace).Return(data.MockUId, nil)

	spaceService := service.NewSpaceService(mockRepo, mockFactory)
	token, err := spaceService.CreateSpace(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, data.MockDescription)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(token), 0, "should return valid uid where the length is greater than 0")
}

// FindSpacesWithinLatLng
func TestFindSpacesWithinLatLng_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockFactory := mocks.NewSpaceFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockDescription).Return(data.MockSpace, nil)
	mockRepo.On("Save", data.MockSpace).Return(data.MockUId, nil)

	spaceService := service.NewSpaceService(mockRepo, mockFactory)
	_, err := spaceService.CreateSpace(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, data.MockDescription)
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByLatLng", data.MockLatitude, data.MockLongitude, data.MockDistance).Return([]space.Entity{data.MockSpace}, nil)

	result, err := spaceService.FindSpacesWithinLatLng(data.MockLatitude, data.MockLongitude, data.MockDistance)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return spaces greater than 0")
	spaceDTO := data.MockSpace.ToDTO()
	assert.Equal(t, spaceDTO, result[0])

}

// FindSpacesByUserId
func TestFindSpacesByUserId_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockFactory := mocks.NewSpaceFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockDescription).Return(data.MockSpace, nil)
	mockRepo.On("Save", data.MockSpace).Return(data.MockUId, nil)

	spaceService := service.NewSpaceService(mockRepo, mockFactory)
	_, err := spaceService.CreateSpace(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, data.MockDescription)
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByUserId", data.MockLenderId).Return([]space.Entity{data.MockSpace}, nil)

	result, err := spaceService.FindSpacesByUserId(data.MockLenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return spaces greater than 0")
	spaceDTO := data.MockSpace.ToDTO()
	assert.Equal(t, spaceDTO, result[0])
}

// FindSpaceById
func TestFindSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockFactory := mocks.NewSpaceFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockDescription).Return(data.MockSpace, nil)
	mockRepo.On("Save", data.MockSpace).Return(data.MockUId, nil)

	spaceService := service.NewSpaceService(mockRepo, mockFactory)
	_, err := spaceService.CreateSpace(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, data.MockDescription)
	assert.Nil(t, err, "should not throw error")

	mockRepo.On("FindOneById", data.MockUId).Return(data.MockSpace, nil)

	result, err := spaceService.FindSpaceById(data.MockUId)
	spaceDTO := data.MockSpace.ToDTO()
	assert.Equal(t, spaceDTO, result, "should return a space which matches the uid")
	assert.Nil(t, err, "should not throw error")
}

// RemoveSpaceById
func TestRemoveSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockFactory := mocks.NewSpaceFactory(t)
	mockFactory.On("New", data.MockTitle, data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockDescription).Return(data.MockSpace, nil)
	mockRepo.On("Save", data.MockSpace).Return(data.MockUId, nil)
	mockRepo.On("Delete", data.MockUId).Return(data.MockUId, nil)

	spaceService := service.NewSpaceService(mockRepo, mockFactory)
	_, err := spaceService.CreateSpace(data.MockLenderId, data.MockStreetAddress, data.MockLatitude, data.MockLongitude, data.MockImageUrls, data.MockTitle, data.MockDescription)
	assert.Nil(t, err, "should not throw error")

	uid, err := spaceService.RemoveSpaceById(data.MockUId)
	assert.Greater(t, len(uid), 0, "greater than 0")
	assert.Nil(t, err, "should not throw error")
}
