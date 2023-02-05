package test

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
	"github.com/kokiebisu/rental-storage/service-user/test/data"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, *errors.CustomError) {
	data.MockUserRepo = mocks.NewUserRepository(t)
	data.MockEventSender = mocks.NewEventSender(t)
	data.MockUserRepo.On("Save", data.MockUser).Return(data.MockUId, nil)

	data.MockEventSender.On("UserCreated", data.MockUser.ToDTO()).Return(nil)
	data.UserService = service.NewUserService(data.MockUserRepo, data.MockEventSender)
	token, err := data.UserService.CreateUser(data.MockUId, data.MockEmailAddress, data.MockFirstName, data.MockLastName, data.MockPassword, []item.DTO{}, data.MockTimeString)
	return token, err
}

// CreateUser
func TestCreateUserSuccess(t *testing.T) {
	uid, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")
	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
}

// RemoveById
func TestRemoveByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	assert.Nil(t, err, "should not throw error")

	data.MockUserRepo.On("Delete", data.MockUId).Return(nil)

	err = data.UserService.RemoveById(data.MockUId)
	assert.Nil(t, err, "should not throw error")
}

// FindById
func TestFindByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	data.MockUserRepo.On("FindOneById", data.MockUId).Return(data.MockUser, nil)

	result, err := data.UserService.FindById(data.MockUId)
	assert.Equal(t, data.MockUser, result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}

// FindByEmail
func TestFindByEmail(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	data.MockUserRepo.On("FindOneByEmail", data.MockEmailAddress).Return(data.MockUser, nil)

	result, err := data.UserService.FindByEmail(data.MockEmailAddress)
	assert.Equal(t, data.MockUser, result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}
