package client

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

var (
	ddc    *dynamodb.Client
	logger *zap.Logger
)

func GetDynamoDBClient() (*dynamodb.Client, *customerror.CustomError) {
	var err *customerror.CustomError
	if ddc != nil {
		return ddc, nil
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}
	if env == "test" {
		// Development mode
		ddc, err = getDynamoDBDockerClient()

	} else {
		// Production mode
		ddc, err = getDynamoDBClient()
	}
	return ddc, err
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
		logger, err = getLoggerDevelopmentClient()
		return logger, err
	} else {
		// Production mode
		logger, err = getLoggerClient()
		return logger, err
	}
}
