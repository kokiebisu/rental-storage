package port

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type EventSender interface {
	UserCreated(user domain.UserDTO) *errors.CustomError
}
