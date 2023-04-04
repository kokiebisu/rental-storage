package port

import (
	"time"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type TokenStore interface {
	SetAccessToken(userID string, token string, expires time.Duration) *customerror.CustomError
	SetRefreshToken(userID string, token string, expires time.Duration) *customerror.CustomError
	DeleteTokens(userID string) *customerror.CustomError
	GetTokens(userID string) (map[string]string, *customerror.CustomError)
}
