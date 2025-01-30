package app

import (
	"log"
	"net/rpc"

	"github.com/fydhfzh/ecommerce-go-application/src/user-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/repository"
)

func StartApp() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	conn, err := db.ConnectDB(&config.DbConfig)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	// instantiate repository
	userRepository := repository.NewUserRepository(conn)

	rpcServer := RPCServer{
		userRepository: userRepository,
	}

	err = rpc.Register(&rpcServer)
	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	err = rpcListen(config.ApplicationConfig)
	if err != nil {
		log.Fatal(err)
	}
}
