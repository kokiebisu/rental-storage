package controller

import (
	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/store"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/client"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

type Controller struct {
	adptr port.Controller
}

func New() port.Controller {
	rc, err := client.GetStoreClient()
	if err != nil {
		panic(err)
	}
	store := store.NewTokenStore(rc)
	cs := service.NewCryptoService()
	ts := service.NewTokenService()
	as := service.NewAuthenticationService(ts, cs, store)
	adptr := adapter.NewApiGatewayAdapter(as)
	return Controller{
		adptr,
	}
}

func (c Controller) SignIn(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.SignIn(event)
}

func (c Controller) SignUp(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.SignUp(event)
}

func (c Controller) Verify(event interface{}) (interface{}, *customerror.CustomError) {
	return c.adptr.Verify(event)
}
