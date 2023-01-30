package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService interface {
	AddItem(name string, imageUrls []string, ownerId string, listingId string) *errors.CustomError
}

type UserService interface {
	CreateUser(uid string, emailAddress string, firstName string, lastName string, password string, items []item.DTO, createdAt string) (string, *errors.CustomError)
	RemoveById(uid string) *errors.CustomError
	FindById(uid string) (user.Entity, *errors.CustomError)
	FindByEmail(emailAddress string) (user.Entity, *errors.CustomError)
}
