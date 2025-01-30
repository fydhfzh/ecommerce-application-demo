package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig      AppConfig
	ServicesConfig ServicesConfig
}

type AppConfig struct {
	Port uint16
}

type ServicesConfig struct {
	UserServiceConfig           UserServiceConfig
	AuthServiceConfig           AuthServiceConfig
	LoggerServiceConfig         LoggerServiceConfig
	ProductCommandServiceConfig ProductCommandServiceConfig
	ProductQueryServiceConfig   ProductQueryServiceConfig
}

type UserServiceConfig struct {
	Host    string
	RpcPort uint
}

type AuthServiceConfig struct {
	Host     string
	GRPCPort uint
}

type LoggerServiceConfig struct {
	Host string
	Port uint
}

type ProductCommandServiceConfig struct {
	Host     string
	GRPCPort uint
}

type ProductQueryServiceConfig struct {
	Host    string
	RPCPort uint
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
		Port: viper.GetUint16("application.port"),
	}

	servicesConfig := ServicesConfig{
		UserServiceConfig: UserServiceConfig{
			Host:    viper.GetString("services.user-service.host"),
			RpcPort: viper.GetUint("services.user-service.rpcPort"),
		},
		AuthServiceConfig: AuthServiceConfig{
			Host:     viper.GetString("services.auth-service.host"),
			GRPCPort: viper.GetUint("services.auth-service.gRpcPort"),
		},
		LoggerServiceConfig: LoggerServiceConfig{
			Host: viper.GetString("services.logger-service.host"),
			Port: viper.GetUint("services.logger-service.port"),
		},
		ProductCommandServiceConfig: ProductCommandServiceConfig{
			Host:     viper.GetString("services.product-command-service.host"),
			GRPCPort: viper.GetUint("services.product-command-service.gRpcPort"),
		},
		ProductQueryServiceConfig: ProductQueryServiceConfig{
			Host:    viper.GetString("services.product-query-service.host"),
			RPCPort: viper.GetUint("services.product-query-service.rpcPort"),
		},
	}

	Config = Configuration{
		AppConfig:      appConfig,
		ServicesConfig: servicesConfig,
	}

	return nil
}
