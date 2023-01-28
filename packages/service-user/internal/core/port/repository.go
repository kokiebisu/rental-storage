package port

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemRepository interface {
	Save(domain.Item) *errors.CustomError
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (domain.Item, *errors.CustomError)
}

type UserRepository interface {
	Setup() *errors.CustomError
	Save(domain.User) (string, *errors.CustomError)
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (domain.User, *errors.CustomError)
	FindOneByEmail(emailAddress string) (domain.User, *errors.CustomError)
}
