package port

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type AuthenticationService interface {
	SignIn(emailAddress string, password string) (string, *customerror.CustomError)
	SignUp(emailAddress string, firstName string, lastName string, password string) (string, *customerror.CustomError)
	Verify(authenticationToken string) (string, *customerror.CustomError)
}

type TokenService interface {
	GenerateToken(uid string) (string, *customerror.CustomError)
	VerifyToken(tokenString string) (*domain.Claims, *customerror.CustomError)
}

type CryptoService interface {
	HashPassword(password string) (string, *customerror.CustomError)
	VerifyPassword(hashedPassword string, plainPassword string) (bool, *customerror.CustomError)
}
