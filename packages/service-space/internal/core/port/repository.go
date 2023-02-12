package port

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceRepository interface {
	Setup() *customerror.CustomError
	Save(space.Entity) (string, *customerror.CustomError)
	Delete(uid string) (string, *customerror.CustomError)
	FindOneById(uid string) (space.Entity, *customerror.CustomError)
	FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]space.Entity, *customerror.CustomError)
	FindManyByUserId(userId string) ([]space.Entity, *customerror.CustomError)
}
