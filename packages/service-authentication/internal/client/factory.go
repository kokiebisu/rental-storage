package client

import (
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	rc     *redis.Client
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

func GetStoreClient() (*redis.Client, *customerror.CustomError) {
	var err *customerror.CustomError
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		rc, err = getRedisDevelopmentClient()
	} else {
		// Production mode
		rc, err = getRedisClient()
	}
	return rc, err
}
