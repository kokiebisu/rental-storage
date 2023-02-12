package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

type UserFactory interface {
	New(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) user.Entity
}

type ItemFactory interface {
	New(name string, ownerId string, spaceId string) item.Entity
}
