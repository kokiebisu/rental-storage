package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestFindSpaces_Success(t *testing.T) {
	mockSpace := data.MockSpace
	s := service.NewSpaceService(data.SpaceRepository)
	uid := uuid.New().String()
	id, err := s.CreateSpace(uid, mockSpace.LenderId, mockSpace.Location, mockSpace.ImageUrls, mockSpace.Title, mockSpace.Description, mockSpace.CreatedAt, mockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
	uid = uuid.New().String()
	id, err = s.CreateSpace(uid, mockSpace.LenderId, mockSpace.Location, mockSpace.ImageUrls, mockSpace.Title, mockSpace.Description, mockSpace.CreatedAt, mockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
	spaces, err := s.FindSpacesByUserId(mockSpace.LenderId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(spaces), 0, "should return 2 spaces")
}
