package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
func InitializeLogger() {
	logger, _ = zap.NewProduction()
}