package port

import (
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/domain"
	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
)

type AuthorizerService interface {
	Authorize(token string) (domain.Claim, *customerror.CustomError)
}
