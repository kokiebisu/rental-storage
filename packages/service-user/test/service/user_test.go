package test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	factory         = &user.Factory{}
	uid             = faker.UUIDDigit()
	emailAddress    = faker.Email()
	firstName       = faker.FirstName()
	lastName        = faker.LastName()
	password        = faker.Password()
	timeString      = faker.TimeString()
	u               = factory.New(uid, firstName, lastName, emailAddress, password, []item.Entity{}, timeString)
	s               *service.UserService
	mockRepo        *mocks.UserRepository
	mockEventSender *mocks.EventSender
	mockFactory     *mocks.UserFactory
)

func setupTest(t *testing.T) (string, *errors.CustomError) {
	mockRepo = mocks.NewUserRepository(t)
	mockEventSender = mocks.NewEventSender(t)
	mockFactory = mocks.NewUserFactory(t)
	mockRepo.On("Save", u).Return(uid, nil)

	mockEventSender.On("UserCreated", u.ToDTO()).Return(nil)
	mockFactory.On("New", uid, firstName, lastName, emailAddress, password, []item.Entity{}, timeString).Return(u)
	s = service.NewUserService(mockRepo, mockEventSender, mockFactory)
	token, err := s.CreateUser(uid, emailAddress, firstName, lastName, password, []item.DTO{}, timeString)
	return token, err
}

// CreateUser
func TestCreateUserSuccess(t *testing.T) {
	uid, err := setupTest(t)
	assert.Greater(t, len(uid), 0, "should return valid uid where the length is greater than 0")
	assert.Nil(t, err, "should not throw error")
}

// RemoveById
func TestRemoveByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	mockRepo.On("Delete", uid).Return(nil)

	err = s.RemoveById(uid)
	assert.Nil(t, err, "should not throw error")
}

// FindById
func TestFindByIdSuccess(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	mockRepo.On("FindOneById", uid).Return(u, nil)

	result, err := s.FindById(uid)
	assert.Equal(t, u, result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}

// FindByEmail
func TestFindByEmail(t *testing.T) {
	_, err := setupTest(t)
	if err != nil {
		panic("setupTest failed")
	}
	mockRepo.On("FindOneByEmail", emailAddress).Return(u, nil)

	result, err := s.FindByEmail(emailAddress)
	assert.Equal(t, u, result, "user should be found")
	assert.Nil(t, err, "should not throw error")
}
