package client

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

func getLoggerClient() (*zap.Logger, *customerror.CustomError) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	return logger, nil
}

func getLoggerDevelopmentClient() (*zap.Logger, *customerror.CustomError) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, customerror.ErrorHandler.LoggerConfigurationError(err)
	}
	return logger, nil
}
