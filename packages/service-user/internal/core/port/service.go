package port

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService interface {
	AddItem(name string, imageUrls []string, ownerId string, listingId string) *errors.CustomError
}

type UserService interface {
	CreateUser(emailAddress string, firstName string, lastName string, password string) (string, *errors.CustomError)
	RemoveById(uid string) *errors.CustomError
	FindById(uid string) (domain.User, *errors.CustomError)
	FindByEmail(emailAddress string) (domain.User, *errors.CustomError)
}
