package unit

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
	"github.com/kokiebisu/rental-storage/service-user/test/data"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, *customerror.CustomError) {
	data.MockUserRepo = mocks.NewUserRepository(t)
	data.MockEventSender = mocks.NewEventSender(t)
	data.MockUserRepo.On("Save", data.MockUser.ToEntity()).Return(data.MockUser.UId, nil)

	data.MockEventSender.On("UserCreated", data.MockUser).Return(nil)
	data.UserService = service.NewUserService(data.MockUserRepo, data.MockEventSender)
	token, err := data.UserService.CreateUser(data.MockUser.UId, data.MockUser.EmailAddress, data.MockUser.FirstName, data.MockUser.LastName, data.MockUser.Password, data.MockUser.Items, data.MockUser.CreatedAt, data.MockUser.UpdatedAt)
	return token, err
}

// CreateUser
func TestCreateUser_Success(t *testing.T) {
	uid, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
}

// RemoveById
func TestRemoveById_Success(t *testing.T) {
	_, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")

	data.MockUserRepo.On("Delete", data.MockUser.UId).Return(data.MockUser.UId, nil)

	uid, err := data.UserService.RemoveById(data.MockUser.UId)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(uid), 0, "greater than 0")
}

// FindById
func TestFindById_Success(t *testing.T) {
	_, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")
	data.MockUserRepo.On("FindById", data.MockUser.UId).Return(data.MockUser.ToEntity(), nil)

	result, err := data.UserService.FindById(data.MockUser.UId)
	assert.Equal(t, data.MockUser.ToEntity(), result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}

// FindByEmail
func TestFindByEmail_Success(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	data.MockUserRepo.On("FindByEmail", data.MockUser.EmailAddress).Return(data.MockUser.ToEntity(), nil)

	result, err := data.UserService.FindByEmail(data.MockUser.EmailAddress)
	assert.Equal(t, data.MockUser.ToEntity(), result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}
