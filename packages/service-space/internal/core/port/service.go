package port

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceService interface {
	FindSpaceById(uid string) (space.DTO, *customerror.CustomError)
	FindSpacesByUserId(userId string) ([]space.DTO, *customerror.CustomError)
	CreateSpace(uid string, lenderId string, location location.DTO, imageUrls []string, title string, description string, createdAt string, updatedAt string) (string, *customerror.CustomError)
	DeleteSpaceById(uid string) (string, *customerror.CustomError)
}
