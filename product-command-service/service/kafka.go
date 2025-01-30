package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/dto"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaProducerService struct {
	producer *kafka.Writer
}

func NewKafkaProducerService() *KafkaProducerService {
	kafkaConfig := common.Config.KafkaConfig

	topicName := kafkaConfig.TopicName
	port := kafkaConfig.Port

	producer := &kafka.Writer{
		Addr:         kafka.TCP(fmt.Sprintf("kafka:%d", port)),
		Topic:        topicName,
		Async:        true,
		RequiredAcks: kafka.RequireOne,
		BatchSize:    100,
		BatchTimeout: time.Second,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	return &KafkaProducerService{
		producer: producer,
	}
}

func (k *KafkaProducerService) SendProductEvent(event dto.ProductEvent) error {
	event.Payload.ID = uuid.UUID{}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = k.producer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProducerService) Close() {
	err := k.producer.Close()
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("Kafka producer closed successfully")
}
