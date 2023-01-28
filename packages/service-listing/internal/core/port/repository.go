package port

import (
	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingRepository interface {
	Setup() *errors.CustomError
	Save(domain.Listing) (string, *errors.CustomError)
	Delete(uid string) *errors.CustomError
	FindOneById(uid string) (domain.Listing, *errors.CustomError)
	FindManyByLatLng(latitude float32, longitude float32, distance int32) ([]domain.Listing, *errors.CustomError)
}
