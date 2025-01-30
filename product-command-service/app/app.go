package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/repository"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/service"
)

func StartApp() {
	err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// instantiate repository
	productRepository := repository.NewProductRepository(db.DB)

	// insantiate service
	kafkaProducerService := service.NewKafkaProducerService()
	defer kafkaProducerService.Close()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go gRPCListen(productRepository, kafkaProducerService)

	<-shutdown

	log.Println("Closing kafka connection")
	kafkaProducerService.Close()
}
