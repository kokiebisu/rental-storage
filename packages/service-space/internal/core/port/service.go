package port

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceService interface {
	FindSpacesWithinLatLng(latitude float64, longitude float64, distance int32) ([]space.DTO, *customerror.CustomError)
	FindSpaceById(uid string) (space.DTO, *customerror.CustomError)
	FindSpacesByUserId(userId string) ([]space.DTO, *customerror.CustomError)
	CreateSpace(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string) (string, *customerror.CustomError)
	RemoveSpaceById(uid string) (string, *customerror.CustomError)
}
