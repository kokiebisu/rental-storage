package port

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingRepository interface {
	Setup() *customerror.CustomError
	Save(listing.Entity) (string, *customerror.CustomError)
	Delete(uid string) *customerror.CustomError
	FindOneById(uid string) (listing.Entity, *customerror.CustomError)
	FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]listing.Entity, *customerror.CustomError)
}
