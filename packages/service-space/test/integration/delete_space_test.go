package integration

import (
	"log"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-space/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSpaceById_Success(t *testing.T) {
	mockSpace := data.MockSpace
	s := service.NewSpaceService(data.SpaceRepository)
	id, err := s.CreateSpace(mockSpace.UId, mockSpace.LenderId, mockSpace.Location, mockSpace.ImageUrls, mockSpace.Title, mockSpace.Description, mockSpace.CreatedAt, mockSpace.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
	uid, err := s.DeleteSpaceById(id)
	if err != nil {
		log.Fatal(err)
	}
	assert.Greater(t, len(uid), 0, "should return valid uid where the lenth is greater than 0")
}
