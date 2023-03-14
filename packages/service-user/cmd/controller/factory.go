package controller

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type Controller struct {
	adptr port.Controller
}

func New() (port.Controller, *customerror.CustomError) {
	db, err := client.GetPostgresClient()
	if err != nil {
		return Controller{}, err
	}
	repo := repository.NewUserRepository(db)
	err = repo.Setup()
	if err != nil {
		return Controller{}, err
	}
	kc, err := client.GetKinesisClient()
	if err != nil {
		return Controller{}, err
	}
	publisher := publisher.NewUserPublisher(kc)
	service := service.NewUserService(repo, publisher)
	adptr := adapter.NewApiGatewayAdapter(service, publisher)
	return Controller{
		adptr,
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
