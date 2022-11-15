package service

import (
	"errors"
	"fmt"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type UserServiceImpl struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(emailAddress string, firstName string, lastName string, password string) error {
	user := domain.CreateUser(firstName, lastName, emailAddress, password)
	err := s.userRepository.Save(user)
	if err != nil {
		return errors.New("failed to create via repository")
	}
	return nil
}

func (s *UserServiceImpl) RemoveById(uid string) error {
	err := s.userRepository.Delete(uid)
	if err != nil {
		return errors.New("failed to remove via repository")
	}
	return nil
}

func (s *UserServiceImpl) FindById(uid string) (domain.User, error) {
	user, err := s.userRepository.FindOneById(uid)
	if err != nil {
		return domain.User{}, errors.New("failed to find via repository")
	}
	return user, nil
}

func (s *UserServiceImpl) FindByEmail(emailAddress string) (domain.User, error) {
	user, err := s.userRepository.FindOneByEmail(emailAddress)
	if err != nil {
		fmt.Println("ENTERED2", err)
		return domain.User{}, errors.New("failed to find via repository")
	}
	return user, nil
}
