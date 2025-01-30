package logging

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	config := zap.Config{
		OutputPaths: []string{"stdout", "logs/order-service.log"},
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("error building zap logger: %v", err)
	}

	return logger.Sugar()
}
