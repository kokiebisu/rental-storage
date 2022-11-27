package port

import domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"

type ListingRepository interface {
	Setup() error
	Save(domain.Listing) (string, error)
	Delete(uid string) error
	FindOneById(uid string) (domain.Listing, error)
	FindManyByLatLng(latitude float32, longitude float32, distance int32) ([]domain.Listing, error)
}
