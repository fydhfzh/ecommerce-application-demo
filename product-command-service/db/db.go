package db

import (
	"fmt"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() error {
	dbConfig := common.Config.DbConfig

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	rawDB, err := DB.DB()
	if err != nil {
		return err
	}

	err = rawDB.Ping()

	if err != nil {
		return err
	}

	DB.AutoMigrate(&model.Product{})

	return nil
}
