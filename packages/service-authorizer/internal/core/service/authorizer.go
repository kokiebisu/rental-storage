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

func (s *AuthorizerService) Authorize(encodedPayload []byte) (domain.Claim, *customerror.CustomError) {
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
