package app

import (
	"fmt"
	"log"

	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/db"
)

func StartApp() {
	err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := common.Config.AppConfig

	err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", appConfig.Port)))
}
