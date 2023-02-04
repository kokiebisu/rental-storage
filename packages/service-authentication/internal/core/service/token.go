package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

// GenerateToken returns a unique jwt token based on the provided email string
func (s *TokenService) GenerateToken(uid string) (string, *errors.CustomError) {
	claims := domain.Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.ErrorHandler.SignTokenError(err)
	}
	return tokenString, nil
}

// verifies the jwt token
func (s *TokenService) VerifyToken(tokenString string) (*domain.Claims, *errors.CustomError) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, errors.ErrorHandler.ClaimParseError(err)
	}
	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, errors.ErrorHandler.ClaimCastError(err)
	}
	if claims.UId == "" {
		return nil, errors.ErrorHandler.ClaimUidEmptyError(err)
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.ErrorHandler.ClaimExpiredError(err)
	}
	return claims, nil
}
