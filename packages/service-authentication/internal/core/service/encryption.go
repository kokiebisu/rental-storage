package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
)

type EncryptionService struct{}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (s *EncryptionService) SignIn(emailAddress string, password string) (string, error) {
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", os.Getenv("SERVICE_API_ENDPOINT"), emailAddress)
	// check if the email address exists in the user db
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return "", errors.New("failed to send request to user service endpoint")
	}
	user := &domain.User{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}

	matched, err := s.verifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("password didn't match")
	}
	response := &port.GenerateJWTTokenPayload{
		UId: user.Uid,
	}
	token, err := s.generateJWTToken(response)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *EncryptionService) SignUp(emailAddress string, firstName string, lastName string, password string) (string, error) {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("cannot hash password")
	}

	updatedUser := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{
		EmailAddress: emailAddress,
		Password:     string(hash),
		FirstName:    firstName,
		LastName:     lastName,
	}

	encodedUpdatedUser, err := json.Marshal(&updatedUser)
	if err != nil {
		return "", err
	}
	userEndpoint := fmt.Sprintf("%s/users", os.Getenv("SERVICE_API_ENDPOINT"))

	resp, err := http.Post(userEndpoint, "application/json", bytes.NewBuffer(encodedUpdatedUser))
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 500 {
		payload := struct {
			Message string `json:"message"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return "", err
		}
		encodedMessage, err := json.Marshal(payload)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(encodedMessage))
	}
	response := &port.GenerateJWTTokenPayload{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	token, err := s.generateJWTToken(response)
	if err != nil {
		return "", err
	}

	tokenPayload := map[string]string{
		"AuthorizationToken": token,
	}
	encoded, err := json.Marshal(tokenPayload)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (s *EncryptionService) Verify(authorizationToken string) (string, error) {
	claims, err := s.verifyJWT(authorizationToken)
	if err != nil {
		return "", err
	}

	encoded, err := json.Marshal(claims)
	if err != nil {
		log.Fatal(err)
	}
	return string(encoded), nil
}

// GenerateToken returns a unique token based on the provided email string
func (s *EncryptionService) generateJWTToken(payload *port.GenerateJWTTokenPayload) (string, error) {
	claims := domain.Claims{
		payload.UId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *EncryptionService) verifyJWT(tokenString string) (*domain.Claims, error) {
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

func (s *EncryptionService) verifyPassword(hashedPassword string, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}
