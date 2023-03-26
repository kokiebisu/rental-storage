package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/domain"
	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
)

type AuthorizerService struct {
}

func NewAuthorizerService() *AuthorizerService {
	return &AuthorizerService{}
}

// Authorize is a function that verifies the jwt token
// and returns the claim if the token is valid
func (s *AuthorizerService) Authorize(token string) (domain.Claim, *customerror.CustomError) {
	body := struct {
		Token string `json:"authorizationToken"`
	}{
		Token: token,
	}
	encodedPayload, err := json.Marshal(&body)
	if err != nil {
		return domain.Claim{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	authenticationEndpoint := fmt.Sprintf("%s/auth/verify", os.Getenv("SERVICE_API_ENDPOINT"))
	// send REST API to verify jwt
	resp, err := http.Post(authenticationEndpoint, "application/json", bytes.NewBuffer(encodedPayload))
	if err != nil {
		return domain.Claim{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	payload := domain.Claim{}
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return domain.Claim{}, nil
	}
	return payload, nil
}
