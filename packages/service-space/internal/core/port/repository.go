package port

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceRepository interface {
	Setup() *customerror.CustomError
	Save(space.Entity) (string, *customerror.CustomError)
	Delete(uid string) (string, *customerror.CustomError)
	FindById(uid string) (space.Entity, *customerror.CustomError)
	FindManyByUserId(userId string) ([]space.Entity, *customerror.CustomError)
}
