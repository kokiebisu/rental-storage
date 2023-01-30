package test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	factory      = domain.Factory{}
	uid          = faker.UUIDDigit()
	emailAddress = faker.Email()
	firstName    = faker.FirstName()
	lastName     = faker.LastName()
	password     = faker.Password()
	timeString   = faker.TimeString()
)

func TestCreateUserSuccess(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)
	u := factory.NewUser(uid, firstName, lastName, emailAddress, password, []item.Entity{}, timeString)
	mockRepo.On("Save", u).Return(uid, nil)
	mockEventSender := mocks.NewEventSender(t)
	mockEventSender.On("UserCreated", u.ToDTO()).Return(nil)
	mockFactory := mocks.NewFactory(t)
	mockFactory.On("NewUser", uid, firstName, lastName, emailAddress, password, []item.Entity{}, timeString).Return(u)
	service := service.NewUserService(mockRepo, mockEventSender, mockFactory)
	token, err := service.CreateUser(uid, emailAddress, firstName, lastName, password, []item.DTO{}, timeString)
	assert.Greater(t, len(token), 0, "should return valid token where the length of the string is greater than 0")
	assert.Nil(t, err, "should not throw error")
}
