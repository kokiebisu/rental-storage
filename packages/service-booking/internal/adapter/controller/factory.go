package controller

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/internal/repository"
)

func New() (*ApiGatewayHandler, *customerror.CustomError) {
	db, err := adapter.GetDBAdapter()
	if err != nil {
		return nil, err
	}
	repo, err := repository.NewBookingRepository(db)
	if err != nil {
		return nil, err
	}
	s := service.NewBookingService(repo)
	return NewApiGatewayHandler(s), nil
}
