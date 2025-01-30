package db

import (
	"fmt"

	"github.com/fydhfzh/ecommerce-go-application/src/order-service/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig common.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
