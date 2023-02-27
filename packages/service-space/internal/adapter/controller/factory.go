package controller

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/kokiebisu/rental-storage/service-space/internal/repository"
)

func New() (*ApiGatewayHandler, *customerror.CustomError) {
	db, err := db.GetInstance()
	if err != nil {
		return nil, err
	}
	repo := repository.NewSpaceRepository(db)
	err = repo.Setup()
	if err != nil {
		return nil, err
	}
	service := service.NewSpaceService(repo)
	return &ApiGatewayHandler{
		service,
	}, nil
}
