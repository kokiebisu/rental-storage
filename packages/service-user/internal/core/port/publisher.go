package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type UserPublisher interface {
	UserCreated(u user.DTO) *customerror.CustomError
}
