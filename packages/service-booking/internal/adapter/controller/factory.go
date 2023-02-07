package controller

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/internal/repository"
)

func New() (*ApiGatewayHandler, *customerror.CustomError) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	repo := repository.NewBookingRepository(db)
	s := service.NewBookingService(repo)
	return NewApiGatewayHandler(s), nil
}
