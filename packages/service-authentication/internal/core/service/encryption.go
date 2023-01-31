package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

type EncryptionService struct{}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (s *EncryptionService) SignIn(emailAddress string, password string) (string, *errors.CustomError) {
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", os.Getenv("SERVICE_API_ENDPOINT"), emailAddress)
	// check if the email address exists in the user db
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return "", errors.ErrorHandler.RequestFailError(err)
	}
	user := &domain.User{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", errors.ErrorHandler.DecodeError("user service endpoint to user domain", err)
	}

	matched, err := s.verifyPassword(user.Password, password)
	if !matched {
		return "", err.(*errors.CustomError)
	}
	response := &port.GenerateJWTTokenPayload{
		UId: user.Uid,
	}
	token, err := s.generateJWTToken(response)
	return token, err.(*errors.CustomError)
}

func (s *EncryptionService) SignUp(emailAddress string, firstName string, lastName string, password string) (string, *errors.CustomError) {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.ErrorHandler.PasswordHashError(err)
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
		return "", errors.ErrorHandler.UnmarshalError("updated user", err)
	}
	if os.Getenv("SERVICE_API_ENDPOINT") == "" {
		return "", errors.ErrorHandler.UndefinedEndPointError(err)
	}
	userEndpoint := fmt.Sprintf("%s/users", os.Getenv("SERVICE_API_ENDPOINT"))
	resp, err := http.Post(userEndpoint, "application/json", bytes.NewBuffer(encodedUpdatedUser))
	if err != nil {
		return "", errors.ErrorHandler.ResponseInvalidError(err)
	}
	if resp.StatusCode == 500 {
		payload := struct {
			StatusCode uint8  `json:"statusCode"`
			Message    string `json:"message"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return "", errors.ErrorHandler.DecodeError("statusCode and message from status code 500 response", err)
		}
		_, err := helper.Stringify(payload)
		return "", err
	}
	response := &port.GenerateJWTTokenPayload{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", errors.ErrorHandler.DecodeError("jwt payload to response", err)
	}

	token, _ := s.generateJWTToken(response)

	return token, nil
}

func (s *EncryptionService) Verify(authorizationToken string) (string, *errors.CustomError) {
	claims, err := s.verifyJWT(authorizationToken)
	if err != nil {
		return "", err
	}
	encoded, err := helper.Stringify(claims)
	return encoded, err
}

// GenerateToken returns a unique token based on the provided email string
func (s *EncryptionService) generateJWTToken(payload *port.GenerateJWTTokenPayload) (string, *errors.CustomError) {
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
		return "", errors.ErrorHandler.SignTokenError(err)
	}
	return tokenString, nil
}

func (s *EncryptionService) verifyJWT(tokenString string) (*domain.Claims, *errors.CustomError) {
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

func (s *EncryptionService) verifyPassword(hashedPassword string, plainPassword string) (bool, *errors.CustomError) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false, errors.ErrorHandler.CompareHashError(err)
	}
	return true, nil
}
