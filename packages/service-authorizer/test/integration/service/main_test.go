package integration

import (
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/service"
	"github.com/kokiebisu/rental-storage/service-authorizer/test/data"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	// var err *customerror.CustomError

	os.Setenv("SERVICE_API_ENDPOINT", "http://localhost:1234")

	data.AuthorizerService = service.NewAuthorizerService()
}
