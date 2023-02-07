package service

import (
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

	"golang.org/x/crypto/bcrypt"
)

type CryptoService struct {
}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

func (s *CryptoService) VerifyPassword(hashedPassword string, plainPassword string) (bool, *customerror.CustomError) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false, customerror.ErrorHandler.CompareHashError(err)
	}
	return true, nil
}

func (s *CryptoService) HashPassword(password string) (string, *customerror.CustomError) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", customerror.ErrorHandler.PasswordGenerationError(err)
	}
	return string(hashed), nil
}
