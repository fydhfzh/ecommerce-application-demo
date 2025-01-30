package common

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ApplicationConfig ApplicationConfig
	DbConfig          DbConfig
}

type ApplicationConfig struct {
	RpcPort uint
}

type DbConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

func LoadConfig() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(wd + "/config/config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	appConfig := ApplicationConfig{
		RpcPort: viper.GetUint("application.rpcPort"),
	}

	dbConfig := DbConfig{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetUint("postgresql.port"),
		User:     viper.GetString("postgresql.user"),
		Password: viper.GetString("postgresql.password"),
		DbName:   viper.GetString("postgresql.dbName"),
	}

	config := Config{
		ApplicationConfig: appConfig,
		DbConfig:          dbConfig,
	}

	return &config, nil
}
