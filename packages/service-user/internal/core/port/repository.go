package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemRepository interface {
	Save(i item.Entity) *errors.CustomError
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (item.Entity, *errors.CustomError)
}

type UserRepository interface {
	Setup() *errors.CustomError
	Save(u user.Entity) (string, *errors.CustomError)
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (user.Entity, *errors.CustomError)
	FindOneByEmail(emailAddress string) (user.Entity, *errors.CustomError)
}
