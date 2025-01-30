package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig AppConfig
	ESConfig  ESConfig
}

type AppConfig struct {
	RPCPort uint
}

type ESConfig struct {
	Host      string
	Port      uint
	IndexName string
}

var (
	Config Configuration
)

func LoadConfig() error {
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.SetConfigFile(workdir + "/config/config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil
	}

	appConfig := AppConfig{
		RPCPort: viper.GetUint("application.rpcPort"),
	}

	esConfig := ESConfig{
		Host:      viper.GetString("elasticsearch.host"),
		Port:      viper.GetUint("elasticsearch.port"),
		IndexName: viper.GetString("elasticsearch.indexName"),
	}

	Config = Configuration{
		AppConfig: appConfig,
		ESConfig:  esConfig,
	}

	return nil
}
