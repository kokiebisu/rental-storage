package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type EventSender interface {
	UserCreated(u user.DTO) *errors.CustomError
}
