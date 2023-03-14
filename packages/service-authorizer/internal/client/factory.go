package client

import (
	"os"

	"go.uber.org/zap"

	customerror "github.com/kokiebisu/rental-storage/service-authorizer/internal/error"
)

var (
	logger *zap.Logger
)

func GetLoggerClient() (*zap.Logger, *customerror.CustomError) {
	var err *customerror.CustomError
	if logger != nil {
		return logger, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		logger, err = getLoggerDevelopmentClient()
	} else {
		// Production mode
		logger, err = getLoggerClient()
	}
	return logger, err
}
