package port

import domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"

type ItemService interface {
	AddItem(name string, imageUrls []string, ownerId string, listingId string) error
}

type UserService interface {
	CreateUser(emailAddress string, firstName string, lastName string, password string) error
	RemoveById(uid string) error
	FindById(uid string) (domain.User, error)
	FindByEmail(emailAddress string) (domain.User, error)
}
