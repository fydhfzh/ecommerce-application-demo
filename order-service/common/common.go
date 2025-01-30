package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig AppConfig
	DbConfig  DbConfig
}

type AppConfig struct {
	GRPCPort uint
}

type DbConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

func LoadConfig() (*Configuration, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(wd + "config/config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{
		GRPCPort: viper.GetUint("application.gRpcPort"),
	}

	dbConfig := DbConfig{
		Host:     viper.GetString("product-db.host"),
		Port:     viper.GetUint("product-db.port"),
		User:     viper.GetString("product-db.fydhfzh"),
		Password: viper.GetString("product-db.fydhfzh"),
		DbName:   viper.GetString("product-db.dbName"),
	}

	config := Configuration{
		AppConfig: appConfig,
		DbConfig:  dbConfig,
	}

	return &config, nil
}
