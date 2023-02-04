package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-authentication/mocks"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

func TestSignUpSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"uid":"%s"}`, data.MockUId)
		_, err := w.Write([]byte(response))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)
	authenticationService := service.NewAuthenticationService(mockTokenService, mockCryptoService)

	mockCryptoService.On("HashPassword", data.MockPassword).Return("Random", nil)
	mockTokenService.On("GenerateToken", data.MockUId).Return(data.MockToken, nil)
	payload, err := authenticationService.SignUp(data.MockEmailAddress, data.MockFirstName, data.MockLastName, data.MockPassword)
	assert.Greater(t, len(payload), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

func TestSignInSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := fmt.Sprintf(`{"user":{"uid":"%s","firstName":"%s","lastName": "%s","emailAddress":"%s","password":"%s","items":"%s","createdAt":"%s","updatedAt":"%s"}}}`,
			data.MockUId,
			data.MockFirstName,
			data.MockLastName,
			data.MockEmailAddress,
			data.MockPassword,
			"[]",
			data.MockTimeString,
			data.MockTimeString,
		)
		_, err := w.Write([]byte(body))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)

	mockCryptoService.On("VerifyPassword", data.MockPassword, data.MockPassword).Return(true, nil)
	mockTokenService.On("GenerateToken", data.MockUId).Return(data.MockToken, nil)

	tokenService := service.NewAuthenticationService(mockTokenService, mockCryptoService)

	payload, err := tokenService.SignIn(data.MockEmailAddress, data.MockPassword)
	assert.Greater(t, len(payload), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

func TestVerifySuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"uid":"1234-5678-9123"}`))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)
	authenticationService := service.NewAuthenticationService(mockTokenService, mockCryptoService)

	mockTokenService.On("VerifyToken", data.MockToken).Return(data.MockClaims, nil)

	payload, err := authenticationService.Verify(data.MockToken)
	assert.Greater(t, len(payload), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}
