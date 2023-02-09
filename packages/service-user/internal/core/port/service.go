package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService interface {
	AddItem(name string, imageUrls []string, ownerId string, listingId string) *customerror.CustomError
}

type UserService interface {
	CreateUser(uid string, emailAddress string, firstName string, lastName string, password string, items []item.DTO, createdAt string) (string, *customerror.CustomError)
	RemoveById(uid string) (string, *customerror.CustomError)
	FindById(uid string) (user.Entity, *customerror.CustomError)
	FindByEmail(emailAddress string) (user.Entity, *customerror.CustomError)
}
