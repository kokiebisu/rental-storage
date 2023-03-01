package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-user/test/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	mockUser := data.MockUser
	s := service.NewUserService(UserRepo, UserPublisher)
	id, err := s.CreateUser(mockUser.UId, mockUser.EmailAddress, mockUser.FirstName, mockUser.LastName, mockUser.Password, mockUser.Items, mockUser.CreatedAt, mockUser.UpdatedAt)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(id), 0, "should return valid uid where the lenth is greater than 0")
}
