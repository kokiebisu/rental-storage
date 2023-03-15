package integration

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
)

func TestFindSpaces_Success(t *testing.T) {
	s := service.NewSpaceService(data.SpaceRepository)
	id, err := s.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	id, err = s.CreateSpace(data.MockSpaces[1].UId, data.MockSpaces[1].LenderId, data.MockSpaces[1].Location, data.MockSpaces[1].ImageUrls, data.MockSpaces[1].Title, data.MockSpaces[1].Description, data.MockSpaces[1].CreatedAt, data.MockSpaces[1].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	spaces, err := s.FindSpaces(0, 2)
	assert.Nil(t, err, "should not throw error")
	assert.Equal(t, 2, len(spaces), "should return 2 spaces")
}

func TestFindSpacesByUserId_Success(t *testing.T) {
	s := service.NewSpaceService(data.SpaceRepository)

	id, err := s.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	id, err = s.CreateSpace(data.MockSpaces[1].UId, data.MockSpaces[0].LenderId, data.MockSpaces[1].Location, data.MockSpaces[1].ImageUrls, data.MockSpaces[1].Title, data.MockSpaces[1].Description, data.MockSpaces[1].CreatedAt, data.MockSpaces[1].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	spaces, err := s.FindSpacesByUserId(data.MockSpaces[0].LenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(spaces), 0, "should return more than 1 space")
}
