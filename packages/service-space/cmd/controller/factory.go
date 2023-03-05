package controller

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/kokiebisu/rental-storage/service-space/internal/repository"
)

type Controller struct {
	adptr port.Controller
}

func New() (port.Controller, *customerror.CustomError) {
	db, err := adapter.GetDBAdapter()
	if err != nil {
		return Controller{}, err
	}
	repo := repository.NewSpaceRepository(db)
	err = repo.Setup()
	if err != nil {
		return Controller{}, err
	}
	service := service.NewSpaceService(repo)
	adptr, _ := adapter.NewControllerAdapter(service)
	return Controller{
		adptr,
	}, nil
}

func (c Controller) FindSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindSpaceById(event)
}

func (c Controller) FindSpaces(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindSpaces(event)
}

func (c Controller) AddSpace(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.AddSpace(event)
}

func (c Controller) DeleteSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.DeleteSpaceById(event)
}
