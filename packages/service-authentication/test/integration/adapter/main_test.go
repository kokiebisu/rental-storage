package integration

import (
	"log"
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/adapter/store"
	"github.com/kokiebisu/rental-storage/service-authentication/internal/client"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	var err *customerror.CustomError
	// Start a PostgreSQL container
	data.RedisClient, err = client.GetStoreClient()
	if err != nil {
		log.Fatal(err)
	}
	data.TokenStore = store.NewTokenStore(data.RedisClient)
}
