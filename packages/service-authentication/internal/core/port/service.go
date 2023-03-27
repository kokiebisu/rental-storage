package port

import (
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type AuthenticationService interface {
	SignIn(emailAddress string, password string) (map[string]domain.Token, *customerror.CustomError)
	SignUp(emailAddress string, firstName string, lastName string, password string) (map[string]domain.Token, *customerror.CustomError)
	Verify(authenticationToken string) (*domain.Claims, *customerror.CustomError)
}

type TokenService interface {
	GenerateAccessToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError)
	GenerateRefreshToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError)
	VerifyToken(tokenString string) (*domain.Claims, *customerror.CustomError)
}

type CryptoService interface {
	HashPassword(password string) (string, *customerror.CustomError)
	VerifyPassword(hashedPassword string, plainPassword string) (bool, *customerror.CustomError)
}
