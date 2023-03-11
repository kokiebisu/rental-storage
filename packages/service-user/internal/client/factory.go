package client

import (
	"database/sql"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"go.uber.org/zap"

	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

var (
	pc     *sql.DB
	kc     *kinesis.Client
	logger *zap.Logger
)

func GetPostgresClient() (*sql.DB, *customerror.CustomError) {
	var err *customerror.CustomError
	if pc != nil {
		return pc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		pc, err = getPostgresDockerClient()

	} else {
		// Production mode
		pc, err = getPostgresClient()
	}
	return pc, err
}

func GetKinesisClient() (*kinesis.Client, *customerror.CustomError) {
	var err *customerror.CustomError
	if kc != nil {
		return kc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		kc, err = getKinesisDockerClient()
	} else {
		// Production mode
		kc, err = getKinesisClient()
	}
	return kc, err
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
