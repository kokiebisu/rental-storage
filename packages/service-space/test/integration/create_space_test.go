package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateSpace_Success(t *testing.T) {
	mockSpace := data.MockSpaces[0]
	service := service.NewSpaceService(data.SpaceRepository)
	id, err := service.CreateSpace(mockSpace.UId, mockSpace.LenderId, mockSpace.Location, mockSpace.ImageUrls, mockSpace.Title, mockSpace.Description, mockSpace.CreatedAt, mockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
}
