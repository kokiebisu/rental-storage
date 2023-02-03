package controller

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/service"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
	"github.com/kokiebisu/rental-storage/service-listing/internal/repository"
)

func New() (*ApiGatewayHandler, *errors.CustomError) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	repo := repository.NewListingRepository(db)
	err = repo.Setup()
	if err != nil {
		return nil, err
	}
	factory := &listing.Factory{}
	service := service.NewListingService(repo, factory)
	return &ApiGatewayHandler{
		service: service,
	}, nil
}
