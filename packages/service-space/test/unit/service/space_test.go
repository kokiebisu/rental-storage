package unit

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
	mockRepo.On("Save", data.MockSpaces[0].ToEntity()).Return(data.MockSpaces[0].UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	token, err := spaceService.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(token), 0, "should return valid uid where the length is greater than 0")
}

// FindAllSpaces
func TestFindSpaces_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpaces[0].ToEntity()).Return(data.MockSpaces[0].UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindMany", 0, 10).Return([]space.Entity{data.MockSpaces[0].ToEntity()}, nil)

	result, err := spaceService.FindSpaces(0, 10)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return spaces greater than 0")
	assert.Equal(t, data.MockSpaces[0], result[0])
}

// FindSpacesByUserId
func TestFindSpacesByUserId_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpaces[0].ToEntity()).Return(data.MockSpaces[0].UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	mockRepo.On("FindManyByUserId", data.MockSpaces[0].LenderId).Return([]space.Entity{data.MockSpaces[0].ToEntity()}, nil)

	result, err := spaceService.FindSpacesByUserId(data.MockSpaces[0].LenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(result), 0, "should return spaces greater than 0")
	assert.Equal(t, data.MockSpaces[0], result[0])
}

// FindSpaceById
func TestFindSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpaces[0].ToEntity()).Return(data.MockSpaces[0].UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	mockRepo.On("FindById", data.MockSpaces[0].UId).Return(data.MockSpaces[0].ToEntity(), nil)

	result, err := spaceService.FindSpaceById(data.MockSpaces[0].UId)
	assert.Equal(t, data.MockSpaces[0], result, "should return a space which matches the uid")
	assert.Nil(t, err, "should not throw error")
}

// DeleteSpaceById
func TestDeleteSpaceById_Success(t *testing.T) {
	mockRepo := mocks.NewSpaceRepository(t)
	mockRepo.On("Save", data.MockSpaces[0].ToEntity()).Return(data.MockSpaces[0].UId, nil)
	mockRepo.On("Delete", data.MockSpaces[0].UId).Return(data.MockSpaces[0].UId, nil)

	spaceService := service.NewSpaceService(mockRepo)
	_, err := spaceService.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")

	uid, err := spaceService.DeleteSpaceById(data.MockSpaces[0].UId)
	assert.Greater(t, len(uid), 0, "greater than 0")
	assert.Nil(t, err, "should not throw error")
}
