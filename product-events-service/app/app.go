package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/repository"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/service"
)

func StartApp() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// instantiate repository
	esRepository := repository.NewESRepository(db.Client)
	err = esRepository.CreateProductIndex()
	if err != nil {
		log.Fatal(err)
	}

	// instantiate service
	kafkaConsumerService := service.NewConsumerService(esRepository)

	log.Println("Consumer is running...")
	go kafkaConsumerService.ReadProductEvent()

	<-sigint

	fmt.Printf("Closing consumer...")
	kafkaConsumerService.Close()
}
