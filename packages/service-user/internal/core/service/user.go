package service

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type UserService struct {
	userRepository port.UserRepository
	userPublisher  port.UserPublisher
}

func NewUserService(userRepository port.UserRepository, userPublisher port.UserPublisher) *UserService {
	return &UserService{
		userRepository,
		userPublisher,
	}
}

func (s *UserService) CreateUser(uid string, emailAddress string, firstName string, lastName string, password string, items []item.DTO, createdAt string, updatedAt string) (string, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("creating user")
	user, err := user.New(uid, firstName, lastName, emailAddress, password, []item.Entity{}, createdAt, updatedAt)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	uid, err = s.userRepository.Save(user)
	if err != nil {
		logger.Error(err.Error())
		return uid, err
	}
	err = s.userPublisher.UserCreated(user.ToDTO())

	// if err != nil {
	// TODO: send to dead letter queue (?)
	// }
	return uid, err
}

func (s *UserService) RemoveById(uid string) (string, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("removing user by id")
	uid, err := s.userRepository.Delete(uid)
	// if err != nil {
	// TODO: user removed event
	// }
	return uid, err
}

func (s *UserService) FindById(uid string) (user.Entity, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("finding user by id")
	u, err := s.userRepository.FindById(uid)
	if err != nil {
		logger.Error(err.Error())
		return user.Entity{}, err
	}
	return u, nil
}

func (s *UserService) FindByEmail(emailAddress string) (user.Entity, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("finding user by email")
	u, err := s.userRepository.FindByEmail(emailAddress)
	if err != nil {
		logger.Error(err.Error())
		return user.Entity{}, err
	}
	return u, nil
}
