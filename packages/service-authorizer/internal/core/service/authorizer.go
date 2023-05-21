package service

import (
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
	authenticationEndpoint := fmt.Sprintf("%s/auth/verify?authorizationToken=%s", os.Getenv("SERVICE_API_ENDPOINT"), token)
	// send REST API to verify jwt
	resp, err := http.Get(authenticationEndpoint)
	if err != nil {
		return domain.Claim{}, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	payload := domain.Claim{}
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return domain.Claim{}, nil
	}
	return payload, nil
}
