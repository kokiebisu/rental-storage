package controller

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/internal/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/repository"
)

type Controller struct {
	adptr port.Controller
}

func New() (port.Controller, *customerror.CustomError) {
	db, err := adapter.GetDBAdapter()
	if err != nil {
		return Controller{}, err
	}
	repo := repository.NewUserRepository(db)
	err = repo.Setup()
	if err != nil {
		return Controller{}, err
	}
	kc, err := adapter.NewKinesisAdapter()
	if err != nil {
		return Controller{}, err
	}
	publisher := publisher.NewUserPublisher(kc)
	service := service.NewUserService(repo, publisher)
	controllerAdapter, _ := adapter.NewControllerAdapter(service)
	return Controller{
		adptr: controllerAdapter,
	}, nil
}

func (c Controller) CreateUser(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.CreateUser(event)
}

func (c Controller) FindUserByEmail(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindUserByEmail(event)
}

func (c Controller) FindUserById(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.FindUserById(event)
}

func (c Controller) RemoveUserById(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.RemoveUserById(event)
}
