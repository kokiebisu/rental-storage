package service

import (
	"errors"
	"fmt"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(emailAddress string, firstName string, lastName string, password string) error {
	user := domain.CreateUser(firstName, lastName, emailAddress, password)
	err := s.userRepository.Save(user)
	if err != nil {
		return errors.New("failed to create via repository")
	}
	return nil
}

func (s *UserService) RemoveById(uid string) error {
	err := s.userRepository.Delete(uid)
	if err != nil {
		return errors.New("failed to remove via repository")
	}
	return nil
}

func (s *UserService) FindById(uid string) (domain.User, error) {
	user, err := s.userRepository.FindOneById(uid)
	if err != nil {
		return domain.User{}, errors.New("failed to find via repository")
	}
	return user, nil
}

func (s *UserService) FindByEmail(emailAddress string) (domain.User, error) {
	user, err := s.userRepository.FindOneByEmail(emailAddress)
	if err != nil {
		fmt.Println("ENTERED2", err)
		return domain.User{}, errors.New("failed to find via repository")
	}
	return user, nil
}
