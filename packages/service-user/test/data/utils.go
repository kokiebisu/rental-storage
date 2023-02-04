package data

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
)

var (
	MockUserRepo    *mocks.UserRepository
	MockEventSender *mocks.EventSender
	UserService     port.UserService
)
