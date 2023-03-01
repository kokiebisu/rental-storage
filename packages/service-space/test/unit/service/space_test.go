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
	mockRepo.On("Save", data.MockSpace.ToEntity()).Return(data.MockSpace.UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	token, err := spaceService.CreateSpace(data.MockSpace.UId, data.MockSpace.LenderId, data.MockSpace.Location, data.MockSpace.ImageUrls, data.MockSpace.Title, data.MockSpace.Description, data.MockSpace.CreatedAt, data.MockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(token), 0, "should return valid uid where the length is greater than 0")
}

// FindSpacesByUserId
func TestFindSpacesByUserId_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpace.ToEntity()).Return(data.MockSpace.UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpace.UId, data.MockSpace.LenderId, data.MockSpace.Location, data.MockSpace.ImageUrls, data.MockSpace.Title, data.MockSpace.Description, data.MockSpace.CreatedAt, data.MockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByUserId", data.MockSpace.LenderId).Return([]space.Entity{data.MockSpace.ToEntity()}, nil)

	result, err := spaceService.FindSpacesByUserId(data.MockSpace.LenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return spaces greater than 0")
	assert.Equal(t, data.MockSpace, result[0])
}

// FindSpaceById
func TestFindSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpace.ToEntity()).Return(data.MockSpace.UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpace.UId, data.MockSpace.LenderId, data.MockSpace.Location, data.MockSpace.ImageUrls, data.MockSpace.Title, data.MockSpace.Description, data.MockSpace.CreatedAt, data.MockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	mockRepo.On("FindById", data.MockSpace.UId).Return(data.MockSpace.ToEntity(), nil)

	result, err := spaceService.FindSpaceById(data.MockSpace.UId)
	assert.Equal(t, data.MockSpace, result, "should return a space which matches the uid")
	assert.Nil(t, err, "should not throw error")
}

// DeleteSpaceById
func TestDeleteSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpace.ToEntity()).Return(data.MockSpace.UId, nil)
	mockRepo.On("Delete", data.MockSpace.UId).Return(data.MockSpace.UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpace.UId, data.MockSpace.LenderId, data.MockSpace.Location, data.MockSpace.ImageUrls, data.MockSpace.Title, data.MockSpace.Description, data.MockSpace.CreatedAt, data.MockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	uid, err := spaceService.DeleteSpaceById(data.MockSpace.UId)
	assert.Greater(t, len(uid), 0, "greater than 0")
	assert.Nil(t, err, "should not throw error")
}
