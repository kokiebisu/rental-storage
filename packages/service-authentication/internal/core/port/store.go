package port

import (
	"time"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type TokenStore interface {
	SetAccessToken(userID, token string, expires time.Duration) *customerror.CustomError
	GetAccessToken(userID string) (string, *customerror.CustomError)
	DeleteAccessToken(userID string) *customerror.CustomError
	SetRefreshToken(userID, token string, expires time.Duration) *customerror.CustomError
	GetRefreshToken(userID string) (string, *customerror.CustomError)
	DeleteRefreshToken(userID string) *customerror.CustomError
	VerifyAccessToken(token string, userID string) *customerror.CustomError
}
