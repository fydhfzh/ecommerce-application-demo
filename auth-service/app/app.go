package app

import (
	"log"

	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/logger"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/service"
)

func StartApp() {
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("error initializing zap logger: %v", err)
	}

	config, err := common.LoadConfig()
	if err != nil {
		logger.Logger.Fatalf("error reading config: %v", err)
	}

	// instantiate service
	jwtService := service.NewJwtService()

	userServiceConfig := config.UserServiceConfig

	authServer := AuthServer{
		jwtService:        jwtService,
		userServiceConfig: userServiceConfig,
	}

	gRpcListen(config.AppConfig.GRpcPort, authServer)

}
