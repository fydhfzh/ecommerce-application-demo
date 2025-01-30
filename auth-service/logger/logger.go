package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

var LOG_PATH = "/app/logs/user-service.log"
var STDOUT = "stdout"

func InitLogger() error {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{STDOUT, LOG_PATH}

	logger, err := config.Build()
	if err != nil {
		return err
	}
	defer logger.Sync()

	Logger = logger.Sugar()

	return nil
}
