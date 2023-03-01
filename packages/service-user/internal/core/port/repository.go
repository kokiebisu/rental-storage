package port

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemRepository interface {
	Save(i item.Entity) (string, *customerror.CustomError)
	Delete(uid string) (string, *customerror.CustomError)
	FindById(uid string) (item.Entity, *customerror.CustomError)
}

type UserRepository interface {
	Setup() *customerror.CustomError
	Save(u user.Entity) (string, *customerror.CustomError)
	Delete(uid string) (string, *customerror.CustomError)
	FindById(uid string) (user.Entity, *customerror.CustomError)
	FindByEmail(emailAddress string) (user.Entity, *customerror.CustomError)
}
