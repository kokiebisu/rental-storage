package adapter

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"service-authentication/pkg/domain"
	"service-authentication/pkg/port"
)

type EncryptionService struct {}

func SetupEncryptionService() (*EncryptionService) {
	return &EncryptionService{}
}

// GenerateToken returns a unique token based on the provided email string
func (s *EncryptionService) GenerateJWTToken(payload *port.GenerateJWTTokenPayload) (string, error) {
    claims := domain.Claims{
        payload.UId,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer: "test",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func (s *EncryptionService) VerifyJWT(tokenString string) (*domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, errors.New("couldn't parse")
	}
	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}
	if claims.UId == "" {
		return nil, errors.New("UID is empty")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("JWT is expired")
	}
	return claims, nil
}

func (s *EncryptionService) VerifyPassword(hashedPassword string, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}