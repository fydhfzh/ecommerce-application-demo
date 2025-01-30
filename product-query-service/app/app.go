package app

import (
	"log"
	"net/rpc"

	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/repository"
)

func StartApp() {
	err := common.LoadConfig()
	if err != nil {
		log.Fatal("error loading configuration")
	}

	err = db.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to elasticsearch")
	}

	// instantiate product repository
	productRepository := repository.NewProductRepository(db.Client)

	// register the RPC server
	rpcServer := NewRPCServer(productRepository)
	err = rpc.Register(rpcServer)
	if err != nil {
		log.Fatal("error register RPC server")
	}
	rpcListen()
}
