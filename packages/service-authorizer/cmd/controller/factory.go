package controller

import (
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
)

type Controller struct {
	adptr port.Controller
}

func New() (port.Controller, *customerror.CustomError) {
	as := service.NewAuthorizerService()
	adptr := adapter.NewApiGatewayAdapter(as)
	return Controller{
		adptr,
	}, nil
}

func (c Controller) Authorize(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.Authorize(event)
}
