package logger

import (
	"log"

	"go.uber.org/zap"
)

var (
	_logger *zap.Logger
)

func Init() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr", "./logs/app.log"}

	var err error
	if _logger, err = cfg.Build(); err != nil {
		log.Panicf("init zap logger with error: %v", err)
	}
}

func Log() *zap.Logger {
	return _logger
}
