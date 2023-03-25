package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

type AuthenticationService struct {
	tokenService  port.TokenService
	cryptoService port.CryptoService
}

func NewAuthenticationService(tokenService port.TokenService, cryptoService port.CryptoService) port.AuthenticationService {
	return &AuthenticationService{
		tokenService,
		cryptoService,
	}
}

// SignIn checks if the email address exists in the user db
// if it does, it checks if the password matches the hashed password
// if it does, it generates a token and returns it
func (s *AuthenticationService) SignIn(emailAddress string, password string) (string, *customerror.CustomError) {
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:1234"
	}
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", endpoint, emailAddress)
	// check if the email address exists in the user db
	resp, err := http.Get(userEndpoint)
	if err != nil {
		return "", customerror.ErrorHandler.RequestFailError(err)
	}
	if resp.StatusCode != 200 {
		return "", customerror.ErrorHandler.RequestInternalError(err)
	}
	payload := struct {
		User user.DTO `json:"user"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return "", customerror.ErrorHandler.DecodeError("user service endpoint to user domain", err)
	}
	matched, err := s.cryptoService.VerifyPassword(payload.User.Password, password)
	if !matched {
		return "", err.(*customerror.CustomError)
	}
	token, err := s.tokenService.GenerateToken(payload.User.UId)
	return token, err.(*customerror.CustomError)
}

// SignUp checks if the email address exists in the user db
// if it does, it returns an error
// if it doesn't, it creates a new user and returns a token
func (s *AuthenticationService) SignUp(emailAddress string, firstName string, lastName string, password string) (string, *customerror.CustomError) {
	// hash password
	hash, err := s.cryptoService.HashPassword(password)
	if err != nil {
		return "", customerror.ErrorHandler.PasswordHashError(err)
	}
	updatedUser := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{
		EmailAddress: emailAddress,
		Password:     hash,
		FirstName:    firstName,
		LastName:     lastName,
	}
	encodedUpdatedUser, err := helper.Stringify(updatedUser)
	if err != nil {
		return "", customerror.ErrorHandler.UnmarshalError("updated user", err)
	}
	baseUrl := os.Getenv("SERVICE_API_ENDPOINT")
	if baseUrl == "" {
		return "", customerror.ErrorHandler.EnvironmentVariableError(nil, "SERVICE_API_ENDPOINT")
	}
	userEndpoint := fmt.Sprintf("%s/users", baseUrl)
	resp, err := helper.SendPostRequest(userEndpoint, encodedUpdatedUser)
	if err != nil {
		return "", customerror.ErrorHandler.ResponseInvalidError(err)
	}
	if resp.StatusCode != 200 {
		payload := struct {
			StatusCode uint8  `json:"statusCode"`
			ErrorCode  string `json:"errorCode"`
			Reason     string `json:"reason"`
			Message    string `json:"message"`
		}{}
		if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return "", customerror.ErrorHandler.DecodeError("statusCode and message from status code 500 response", err)
		}
		internalServiceInfo, err := helper.Stringify(payload)
		return internalServiceInfo, err
	}

	response := &struct {
		UId string `json:"uid"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", customerror.ErrorHandler.DecodeError("jwt payload to response", err)
	}
	return s.tokenService.GenerateToken(response.UId)
}

// Verify checks if the token is valid
func (s *AuthenticationService) Verify(authorizationToken string) (*domain.Claims, *customerror.CustomError) {
	return s.tokenService.VerifyToken(authorizationToken)
}
