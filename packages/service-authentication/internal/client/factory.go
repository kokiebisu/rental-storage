package client

import (
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"go.uber.org/zap"
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
		return logger, err
	} else {
		// Production mode
		logger, err = getLoggerClient()
		return logger, err
	}
}
