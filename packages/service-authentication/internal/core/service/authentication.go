package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
)

type AuthenticationService struct {
	tokenService  port.TokenService
	cryptoService port.CryptoService
}

func NewAuthenticationService(tokenService port.TokenService, cryptoService port.CryptoService) *AuthenticationService {
	return &AuthenticationService{
		tokenService,
		cryptoService,
	}
}

func (s *AuthenticationService) SignIn(emailAddress string, password string) (string, *customerror.CustomError) {
	if os.Getenv("SERVICE_API_ENDPOINT") == "" {
		return "", customerror.ErrorHandler.UndefinedEndPointError(nil)
	}
	userEndpoint := fmt.Sprintf("%s/users/find-by-email?emailAddress=%s", os.Getenv("SERVICE_API_ENDPOINT"), emailAddress)
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
	if os.Getenv("SERVICE_API_ENDPOINT") == "" {
		return "", customerror.ErrorHandler.UndefinedEndPointError(nil)
	}
	userEndpoint := fmt.Sprintf("%s/users", os.Getenv("SERVICE_API_ENDPOINT"))
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
	token, _ := s.tokenService.GenerateToken(response.UId)
	return token, nil
}

func (s *AuthenticationService) Verify(authorizationToken string) (string, *customerror.CustomError) {
	claims, err := s.tokenService.VerifyToken(authorizationToken)
	if err != nil {
		return "", err
	}
	encoded, err := helper.Stringify(claims)
	return encoded, err
}
