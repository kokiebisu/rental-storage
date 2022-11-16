package port

import domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"

type EventSender interface {
	UserCreated(user domain.UserDTO) error
}
