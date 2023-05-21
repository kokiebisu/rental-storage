package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type TokenService struct {
}

func NewTokenService() port.TokenService {
	return &TokenService{}
}

var SECRET_KEY = "SECRET"

// GenerateAccessToken generates an access token
// that expires in 24 hours
func (s *TokenService) GenerateAccessToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError) {
	return s.generateToken(uid, expiresAt)
}

// GenerateRefreshToken generates a refresh token
// that expires in 7 days
// It is used to generate a new access token
func (s *TokenService) GenerateRefreshToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError) {
	return s.generateToken(uid, expiresAt)
}

// GenerateAccessToken generates a token
// If it fails, it returns an error
// If it succeeds, it returns the token
func (s *TokenService) generateToken(uid string, expiresAt time.Duration) (domain.Token, *customerror.CustomError) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expiresAt).Unix()
	claims["authorized"] = true
	claims["uid"] = uid
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", customerror.ErrorHandler.SignTokenError(err)
	}
	return domain.Token(tokenString), nil
}

// VerifyToken verifies the token.
// If it does not match, it returns an error
// If the claims are empty, it returns an error
// If the claims are expired, it returns an error
// If the claims are valid, it returns the claims
func (s *TokenService) VerifyToken(tokenString string) (*domain.Claims, *customerror.CustomError) {
	token, _ := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, customerror.ErrorHandler.ClaimCastError(fmt.Errorf(""))
	}
	if claims.UId == "" {
		return nil, customerror.ErrorHandler.ClaimUIdEmptyError(fmt.Errorf(""))
	}
	timeNow := time.Now().In(time.UTC)
	expTime := time.Unix(claims.ExpiresAt, 0).In(time.UTC)
	if expTime.Before(timeNow) {
		return nil, customerror.ErrorHandler.ClaimExpiredError(fmt.Errorf(""))
	}
	return claims, nil
}
