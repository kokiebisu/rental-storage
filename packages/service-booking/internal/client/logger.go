package client

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLoggerClient() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}
