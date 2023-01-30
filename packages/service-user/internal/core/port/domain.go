package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

type Factory interface {
	NewUser(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) user.Entity
	NewItem(name string, ownerId string, listingId string) item.Entity
}
