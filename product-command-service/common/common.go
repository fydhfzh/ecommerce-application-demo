package common

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppConfig   AppConfig
	DbConfig    DbConfig
	KafkaConfig KafkaConfig
}

type AppConfig struct {
	GRPCPort uint
}

type DbConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	DbName   string
}

type KafkaConfig struct {
	TopicName string
	Port      uint
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
		GRPCPort: viper.GetUint("application.gRpcPort"),
	}

	dbConfig := DbConfig{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetUint("postgresql.port"),
		Username: viper.GetString("postgresql.username"),
		Password: viper.GetString("postgresql.password"),
		DbName:   viper.GetString("postgresql.dbName"),
	}

	kafkaConfig := KafkaConfig{
		TopicName: viper.GetString("kafka.topicName"),
		Port:      viper.GetUint("kafka.port"),
	}

	Config = Configuration{
		AppConfig:   appConfig,
		DbConfig:    dbConfig,
		KafkaConfig: kafkaConfig,
	}

	return nil
}
