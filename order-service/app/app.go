package app

import (
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/logging"
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/repository"
)

func StartApp() {
	logger := logging.NewLogger()

	config, err := common.LoadConfig()
	if err != nil {
		logger.Fatalf("error loading config: %v", err)
	}

	db, err := db.ConnectDB(config.DbConfig)
	if err != nil {
		logger.Fatalf("error connecting to database: %v", err)
	}

	// instantiate repository
	orderRepository := repository.NewRepository(db)

	// instantiate grpc server
	orderServer := OrderServer{
		orderRepository: orderRepository,
	}

	err = grpcListen(config.AppConfig.GRPCPort, orderServer)
	if err != nil {
		logger.Fatalf("error starting grpc server: %v", err)
	}

}
