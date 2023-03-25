package integration

import (
	"log"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/store"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/client"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	var err *customerror.CustomError

	os.Setenv("SERVICE_API_ENDPOINT", "http://localhost:1234")

	data.RedisClient, err = client.GetStoreClient()
	if err != nil {
		log.Fatal(err)
	}
	data.TokenStore = store.NewTokenStore(data.RedisClient)

	data.CryptoService = service.NewCryptoService()
	data.TokenService = service.NewTokenService()
	data.AuthenticationService = service.NewAuthenticationService(data.TokenService, data.CryptoService)
}
