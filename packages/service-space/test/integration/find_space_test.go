package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestFindSpace_Success(t *testing.T) {
	mockSpace := data.MockSpace
	s := service.NewSpaceService(Repo)
	id, err := s.CreateSpace(mockSpace.UId, mockSpace.LenderId, mockSpace.Location, mockSpace.ImageUrls, mockSpace.Title, mockSpace.Description, mockSpace.CreatedAt, mockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	spaceDTO, err := s.FindSpaceById(id)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(spaceDTO.UId), 0, "should return valid uid where the lenth is greater than 0")
}
