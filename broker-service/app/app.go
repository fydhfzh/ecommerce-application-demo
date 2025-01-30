package app

import (
	"fmt"
	"log"

	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/common"
	_ "github.com/fydhfzh/ecommerce-go-application/src/broker-service/docs"
)

func StartApp() {
	err := common.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	appConfig := common.Config.AppConfig

	router := NewRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", appConfig.Port)))
}
