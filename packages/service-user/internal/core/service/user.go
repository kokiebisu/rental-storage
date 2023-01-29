package service

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type UserService struct {
	userRepository port.UserRepository
	eventSender    port.EventSender
}

func NewUserService(userRepository port.UserRepository, eventSender port.EventSender) *UserService {
	return &UserService{
		userRepository: userRepository,
		eventSender:    eventSender,
	}
}

func (s *UserService) CreateUser(emailAddress string, firstName string, lastName string, password string) (string, *errors.CustomError) {
	user := domain.CreateUser(firstName, lastName, emailAddress, password)
	uid, err := s.userRepository.Save(user)
	if err != nil {
		return uid, err
	}
	err = s.eventSender.UserCreated(user.ToDTO())

	// if err != nil {
	// TODO: send to dead letter queue (?)
	// }
	return uid, err
}

func (s *UserService) RemoveById(uid string) *errors.CustomError {
	err := s.userRepository.Delete(uid)
	// if err != nil {
	// TODO: user removed event
	// }
	return err
}

func (s *UserService) FindById(uid string) (domain.User, *errors.CustomError) {
	user, err := s.userRepository.FindOneById(uid)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *UserService) FindByEmail(emailAddress string) (domain.User, *errors.CustomError) {
	user, err := s.userRepository.FindOneByEmail(emailAddress)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
