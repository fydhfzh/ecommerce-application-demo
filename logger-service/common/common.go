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
	Port int
}

type DbConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DbName       string
	DbCollection string
}

var (
	Config Configuration
)

func LoadConfig() error {
	workDir, err := os.Getwd()

	if err != nil {
		return err
	}

	viper.SetConfigFile(workDir + "/config/config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	appConfig := AppConfig{
		Port: viper.GetInt("application.port"),
	}

	dbConfig := DbConfig{
		Host:         viper.GetString("mongodb.host"),
		Port:         viper.GetInt("mongodb.port"),
		Username:     viper.GetString("mongodb.username"),
		Password:     viper.GetString("mongodb.password"),
		DbName:       viper.GetString("mongodb.dbName"),
		DbCollection: viper.GetString("mongodb.dbCollection"),
	}

	Config = Configuration{
		AppConfig: appConfig,
		DbConfig:  dbConfig,
	}

	return nil
}
