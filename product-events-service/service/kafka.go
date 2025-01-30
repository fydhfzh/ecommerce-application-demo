package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/constant"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/repository"
	"github.com/segmentio/kafka-go"
)

type kafkaConsumerService struct {
	consumer     *kafka.Reader
	esRepository repository.ProductRepository
}

type ConsumerService interface {
	ReadProductEvent()
	Close()
}

func NewConsumerService(esRepository repository.ProductRepository) ConsumerService {
	kafkaConfig := common.Config.KafkaConfig

	consumer := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{fmt.Sprintf("%s:%d", kafkaConfig.Host, kafkaConfig.Port)},
			Topic:   kafkaConfig.TopicName,
			GroupID: "product-consumer",
		},
	)

	return &kafkaConsumerService{
		consumer:     consumer,
		esRepository: esRepository,
	}
}

func (k *kafkaConsumerService) ReadProductEvent() {
	for {
		msg, err := k.consumer.FetchMessage(context.Background())
		if errors.Is(err, context.Canceled) {
			log.Println("Server is not responding")
			continue
		}

		var productEvent dto.ProductEvent

		if err := json.Unmarshal(msg.Value, &productEvent); err != nil {
			log.Println("error binding message to product model")
			continue
		}

		switch productEvent.EventType {
		case constant.SaveProduct, constant.UpdateProduct:
			err = k.StoreProductToElasticsearch(productEvent.Payload)
			if err != nil {
				log.Println("error storing message to elasticsearch")
				continue
			}
		default:
			log.Println("event type not found")
		}

		log.Printf("Message: %s | %v", msg.Value, time.Now())
	}
}

func (k *kafkaConsumerService) StoreProductToElasticsearch(product dto.Product) error {
	err := k.esRepository.SaveProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (k *kafkaConsumerService) Close() {
	err := k.consumer.Close()
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Kafka producer closed successfully")
}
