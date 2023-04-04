package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

type AuthenticationService struct {
	tokenService  port.TokenService
	cryptoService port.CryptoService
	tokenStore    port.TokenStore
}

func NewAuthenticationService(tokenService port.TokenService, cryptoService port.CryptoService, tokenStore port.TokenStore) port.AuthenticationService {
	return &AuthenticationService{
		tokenService,
		cryptoService,
		tokenStore,
	}
}

// SignIn checks if the email address exists in the user db
// if it does, it checks if the password matches the hashed password
// if it does, it generates a token and returns it
func (s *AuthenticationService) SignIn(emailAddress string, password string) (map[string]domain.Token, *customerror.CustomError) {
	var e error
	endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:1234"
	}
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", endpoint, emailAddress)
	// check if the email address exists in the user db
	resp, e := http.Get(userEndpoint)
	if e != nil {
		return map[string]domain.Token{}, customerror.ErrorHandler.RequestFailError(e)
	}
	if resp.StatusCode != 200 {
		return map[string]domain.Token{}, customerror.ErrorHandler.RequestInternalError(e)
	}
	payload := struct {
		User user.DTO `json:"user"`
	}{}
	if e = json.NewDecoder(resp.Body).Decode(&payload); e != nil {
		return map[string]domain.Token{}, customerror.ErrorHandler.DecodeError("user service endpoint to user domain", e)
	}
	matched, err := s.cryptoService.VerifyPassword(payload.User.Password, password)
	if !matched {
		return map[string]domain.Token{}, err
	}
	dayDuration := time.Hour * 24

	at, err := s.tokenService.GenerateAccessToken(payload.User.UId, dayDuration)
	if err != nil {
		return map[string]domain.Token{}, err
	}

	// set the access token in the store
	err = s.tokenStore.SetAccessToken(payload.User.UId, string(at), dayDuration)
	if err != nil {
		return map[string]domain.Token{}, err
	}
	rt, err := s.tokenService.GenerateRefreshToken(payload.User.UId, dayDuration*7)
	if err != nil {
		return map[string]domain.Token{}, err
	}
	// set the refresh token in the store
	err = s.tokenStore.SetRefreshToken(payload.User.UId, string(rt), dayDuration*7)
	return map[string]domain.Token{
		"access_token":  at,
		"refresh_token": rt,
	}, err
}

// SignUp checks if the email address exists in the user db
// if it does, it returns an error
// if it doesn't, it creates a new user and returns a token
func (s *AuthenticationService) SignUp(emailAddress string, firstName string, lastName string, password string) (map[string]domain.Token, *customerror.CustomError) {
	// hash password
	hash, err := s.cryptoService.HashPassword(password)
	if err != nil {
		return map[string]domain.Token{}, customerror.ErrorHandler.PasswordHashError(err)
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
		return map[string]domain.Token{}, customerror.ErrorHandler.UnmarshalError("updated user", err)
	}
	baseUrl := os.Getenv("SERVICE_API_ENDPOINT")
	if baseUrl == "" {
		return map[string]domain.Token{}, customerror.ErrorHandler.EnvironmentVariableError(nil, "SERVICE_API_ENDPOINT")
	}
	userEndpoint := fmt.Sprintf("%s/users", baseUrl)
	resp, err := helper.SendPostRequest(userEndpoint, encodedUpdatedUser)
	if err != nil {
		return map[string]domain.Token{}, customerror.ErrorHandler.ResponseInvalidError(err)
	}
	if resp.StatusCode != 200 {
		payload := struct {
			StatusCode uint8  `json:"statusCode"`
			ErrorCode  string `json:"errorCode"`
			Reason     string `json:"reason"`
			Message    string `json:"message"`
		}{}
		if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			return map[string]domain.Token{}, customerror.ErrorHandler.DecodeError("statusCode and message from status code 500 response", err)
		}
		_, err := helper.Stringify(payload)
		return map[string]domain.Token{}, err
	}

	response := &struct {
		UId string `json:"uid"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return map[string]domain.Token{}, customerror.ErrorHandler.DecodeError("jwt payload to response", err)
	}
	dayDuration := time.Hour * 24

	at, err := s.tokenService.GenerateAccessToken(response.UId, dayDuration)
	if err != nil {
		return map[string]domain.Token{}, err
	}
	// set the access token in the store
	err = s.tokenStore.SetAccessToken(response.UId, string(at), dayDuration)
	if err != nil {
		return map[string]domain.Token{}, err
	}
	rt, err := s.tokenService.GenerateRefreshToken(response.UId, dayDuration*7)
	if err != nil {
		return map[string]domain.Token{}, err
	}
	// set the refresh token in the store
	err = s.tokenStore.SetRefreshToken(response.UId, string(rt), dayDuration*7)
	return map[string]domain.Token{
		"access_token":  at,
		"refresh_token": rt,
	}, err
}

// Verify checks if the token is valid
func (s *AuthenticationService) Verify(authorizationToken string) (*domain.Claims, *customerror.CustomError) {
	return s.tokenService.VerifyToken(authorizationToken)
}
