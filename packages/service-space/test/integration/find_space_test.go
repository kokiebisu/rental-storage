package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestFindSpace_Success(t *testing.T) {
	s := service.NewSpaceService(data.SpaceRepository)
	id, err := s.CreateSpace(data.MockSpaces[0].UId, data.MockSpaces[0].LenderId, data.MockSpaces[0].Location, data.MockSpaces[0].ImageUrls, data.MockSpaces[0].Title, data.MockSpaces[0].Description, data.MockSpaces[0].CreatedAt, data.MockSpaces[0].UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")

	spaceDTO, err := s.FindSpaceById(id)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(spaceDTO.UId), 0, "should return valid uid where the lenth is greater than 0")
}
