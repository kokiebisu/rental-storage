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
		// message := fmt.Sprintf("failed to create user in db: %s", err.Error())
		return uid, errors.ErrorHandler.InternalServerError()
	}
	err = s.eventSender.UserCreated(user.ToDTO())
	if err != nil {
		// send to dead letter queue (?)
		// message := fmt.Sprintf("failed to send user created event: %s", err.Error())
		return uid, errors.ErrorHandler.InternalServerError()
	}
	return uid, nil
}

func (s *UserService) RemoveById(uid string) *errors.CustomError {
	err := s.userRepository.Delete(uid)
	if err != nil {
		// message := fmt.Sprintf("failed to remove account from db: %s", err.Error())
		return errors.ErrorHandler.InternalServerError()
	}
	// user removed event
	return nil
}

func (s *UserService) FindById(uid string) (domain.User, *errors.CustomError) {
	user, err := s.userRepository.FindOneById(uid)
	if err != nil {
		// message := fmt.Sprintf("failed to find from db: %s", err.Error())
		return domain.User{}, errors.ErrorHandler.InternalServerError()
	}
	return user, nil
}

func (s *UserService) FindByEmail(emailAddress string) (domain.User, *errors.CustomError) {
	user, err := s.userRepository.FindOneByEmail(emailAddress)
	if err != nil {
		// message := fmt.Sprintf("failed to find from db: %s", err)
		return domain.User{}, errors.ErrorHandler.InternalServerError()
	}
	return user, nil
}
