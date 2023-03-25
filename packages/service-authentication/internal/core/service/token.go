package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type TokenService struct{}

func NewTokenService() port.TokenService {
	return &TokenService{}
}

var SECRET_KEY = "SECRET"

// GenerateToken generates a jwt token for the user based on the uid.
// If it does not match, it returns an error.
// If the claims are empty, it returns an error.
// If the claims are expired, it returns an error.
// If the claims are valid, it returns the claims.
func (s *TokenService) GenerateToken(uid string) (string, *customerror.CustomError) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true
	claims["uid"] = uid
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", customerror.ErrorHandler.SignTokenError(err)
	}
	return tokenString, nil
}

// VerifyToken verifies the token.
// If it does not match, it returns an error
// If the claims are empty, it returns an error
// If the claims are expired, it returns an error
// If the claims are valid, it returns the claims
func (s *TokenService) VerifyToken(tokenString string) (*domain.Claims, *customerror.CustomError) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, customerror.ErrorHandler.ClaimParseError(err)
	}
	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, customerror.ErrorHandler.ClaimCastError(err)
	}
	if claims.UId == "" {
		return nil, customerror.ErrorHandler.ClaimUIdEmptyError(err)
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, customerror.ErrorHandler.ClaimExpiredError(err)
	}
	return claims, nil
}
