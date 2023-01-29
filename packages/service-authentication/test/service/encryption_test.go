package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/domain"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
)

type TestResult struct {
	input    domain.User
	expected interface{}
}

var signUpResults = []TestResult{
	{
		domain.User{
			FirstName:    "Ben",
			LastName:     "Parker",
			EmailAddress: "random@gmail.com",
			Password:     "password123",
		},
		"some token",
	},
}

var s = service.NewEncryptionService()

// for a valid return value.
func TestSignUpSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"uid":"1234-5678-9123"}`))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	for _, test := range signUpResults {
		payload, err := s.SignUp(test.input.EmailAddress, test.input.FirstName, test.input.LastName, test.input.Password)
		assert.Greater(t, len(payload), 0, "token should have a length greater than 0")
		assert.Nil(t, err, "should be no errors")
	}
}
