package data

import (
	"github.com/kokiebisu/rental-storage/service-authorizer/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-authorizer/mocks"
)

var (
	AuthorizerService     port.AuthorizerService
	MockAuthorizerService *mocks.AuthorizerService
)
