package port

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingRepository interface {
	Setup() *errors.CustomError
	Save(listing.Entity) (string, *errors.CustomError)
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (listing.Entity, *errors.CustomError)
	FindManyByLatLng(latitude float32, longitude float32, distance int32) ([]listing.Entity, *errors.CustomError)
}
