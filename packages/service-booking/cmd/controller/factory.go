package controller

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
	"github.com/kokiebisu/rental-storage/service-booking/internal/repository"
)

type Controller struct {
	adptr port.Controller
}

func New() (port.Controller, *customerror.CustomError) {
	db, err := adapter.GetDBAdapter()
	if err != nil {
		return Controller{}, err
	}
	repo, err := repository.NewBookingRepository(db)
	if err != nil {
		return Controller{}, err
	}
	service := service.NewBookingService(repo)
	adptr, _ := adapter.NewControllerAdapter(service)
	return Controller{
		adptr,
	}, nil
}

func (c Controller) CreateBooking(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.CreateBooking(event)
}

func (c Controller) FindBookingById(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindBookingById(event)
}

func (c Controller) FindBookings(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindBookings(event)
}
