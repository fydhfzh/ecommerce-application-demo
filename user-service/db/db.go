package db

import (
	"fmt"

	"github.com/fydhfzh/ecommerce-go-application/src/user-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig *common.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = conn.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	exampleUser := entity.User{
		Email:    "fayyadhhafizh5@gmail.com",
		Password: "test123",
		Fullname: "Fayyadh Hafizh",
		Age:      20,
		Role:     "admin",
	}

	err = exampleUser.HashPassword()
	if err != nil {
		return nil, err
	}

	res := conn.Create(&exampleUser)
	if err := res.Error; err != nil {
		return nil, err
	}

	return conn, nil
}
