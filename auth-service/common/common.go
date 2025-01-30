package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig AppConfig
	UserServiceConfig
}

type AppConfig struct {
	GRpcPort uint
}

type UserServiceConfig struct {
	Host    string
	RpcPort uint
}

func LoadConfig() (*Configuration, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(wd + "/config/config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{
		GRpcPort: viper.GetUint("application.gRpcPort"),
	}

	userServiceConfig := UserServiceConfig{
		Host:    viper.GetString("user-service.host"),
		RpcPort: viper.GetUint("user-service.rpcPort"),
	}

	config := Configuration{
		AppConfig:         appConfig,
		UserServiceConfig: userServiceConfig,
	}

	return &config, err
}
