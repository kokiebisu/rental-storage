package unit

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/kokiebisu/rental-storage/service-authentication/mocks"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

// SignUp
func TestSignUp_Success(t *testing.T) {
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
	mockTokenStore := mocks.NewTokenStore(t)
	authenticationService := service.NewAuthenticationService(mockTokenService, mockCryptoService, mockTokenStore)

	mockTokenStore.On("SetAccessToken", data.MockUId, string(data.MockToken), time.Hour*24).Return(nil)
	mockTokenStore.On("SetRefreshToken", data.MockUId, string(data.MockToken), time.Hour*24*7).Return(nil)
	mockCryptoService.On("HashPassword", data.MockPassword).Return("Random", nil)
	mockTokenService.On("GenerateAccessToken", data.MockUId, time.Hour*24).Return(data.MockToken, nil)
	mockTokenService.On("GenerateRefreshToken", data.MockUId, time.Hour*24*7).Return(data.MockToken, nil)
	payload, err := authenticationService.SignUp(data.MockEmailAddress, data.MockFirstName, data.MockLastName, data.MockPassword)
	assert.Greater(t, len(payload["access_token"]), 0, "token should have a length greater than 0")
	assert.Greater(t, len(payload["refresh_token"]), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

// SignIn
func TestSignIn_Success(t *testing.T) {
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

	mockTokenStore := mocks.NewTokenStore(t)
	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)

	mockTokenStore.On("SetAccessToken", data.MockUId, string(data.MockToken), time.Hour*24).Return(nil)
	mockTokenStore.On("SetRefreshToken", data.MockUId, string(data.MockToken), time.Hour*24*7).Return(nil)
	mockCryptoService.On("VerifyPassword", data.MockPassword, data.MockPassword).Return(true, nil)
	mockTokenService.On("GenerateAccessToken", data.MockUId, time.Hour*24).Return(data.MockToken, nil)
	mockTokenService.On("GenerateRefreshToken", data.MockUId, time.Hour*24*7).Return(data.MockToken, nil)

	authService := service.NewAuthenticationService(mockTokenService, mockCryptoService, mockTokenStore)

	payload, err := authService.SignIn(data.MockEmailAddress, data.MockPassword)
	assert.Greater(t, len(payload["access_token"]), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}

func TestSignIn_Failure_InvalidPassword(t *testing.T) {
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

	mockTokenStore := mocks.NewTokenStore(t)
	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)

	mockCryptoService.On("VerifyPassword", data.MockPassword, "invalid").Return(false, customerror.ErrorHandler.CompareHashError(nil))

	authService := service.NewAuthenticationService(mockTokenService, mockCryptoService, mockTokenStore)

	payload, err := authService.SignIn(data.MockEmailAddress, "invalid")
	assert.NotNil(t, err, "should be an error")
	assert.Equal(t, len(payload), 0, "token should have a length greater than 0")
}

// Verify
func TestVerify_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"uid":"1234-5678-9123"}`))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	mockTokenStore := mocks.NewTokenStore(t)
	mockTokenService := mocks.NewTokenService(t)
	mockCryptoService := mocks.NewCryptoService(t)

	mockTokenService.On("VerifyToken", string(data.MockToken)).Return(data.MockClaims, nil)

	authenticationService := service.NewAuthenticationService(mockTokenService, mockCryptoService, mockTokenStore)

	payload, err := authenticationService.Verify(string(data.MockToken))
	assert.Greater(t, len(payload.UId), 0, "token should have a length greater than 0")
	assert.Nil(t, err, "should be no errors")
}
