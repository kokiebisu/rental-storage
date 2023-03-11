package client

import (
	"database/sql"
	"os"

	"go.uber.org/zap"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

var (
	pc     *sql.DB
	logger *zap.Logger
)

func GetPostgresClient() (*sql.DB, *customerror.CustomError) {
	if pc != nil {
		return pc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		return getPostgresDockerClient()
	} else {
		// Production mode
		return getPostgresClient()
	}
}

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
		logger, err = getLoggerClient()
	} else {
		// Production mode
		logger, err = getLoggerClient()
	}
	return logger, err
}
