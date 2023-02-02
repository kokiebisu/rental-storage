package controller

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/sender"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/service"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/internal/repository"
)

func New() (*ApiGatewayHandler, *errors.CustomError) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	repo := repository.NewUserRepository(db)
	eventSender := sender.New()
	err = repo.Setup()
	factory := &user.Factory{}
	service := service.NewUserService(repo, eventSender, factory)
	return &ApiGatewayHandler{
		service: service,
	}, err
}
