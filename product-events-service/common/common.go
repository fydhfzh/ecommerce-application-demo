package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig   AppConfig
	KafkaConfig KafkaConfig
	ESConfig    ESConfig
}

type AppConfig struct {
	Port uint
}

type KafkaConfig struct {
	Host      string
	Port      uint
	TopicName string
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
		Port: viper.GetUint("application.port"),
	}

	kafkaConfig := KafkaConfig{
		Host:      viper.GetString("kafka.host"),
		Port:      viper.GetUint("kafka.port"),
		TopicName: viper.GetString("kafka.topicName"),
	}

	esConfig := ESConfig{
		Host:      viper.GetString("elasticsearch.host"),
		Port:      viper.GetUint("elasticsearch.port"),
		IndexName: viper.GetString("elasticsearch.indexName"),
	}

	Config = Configuration{
		AppConfig:   appConfig,
		KafkaConfig: kafkaConfig,
		ESConfig:    esConfig,
	}

	return nil
}
