package dto

import (
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/constant"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/model"
)

type ProductEvent struct {
	EventType constant.Method `json:"event_type"`
	ProductId string          `json:"product_id"`
	Timestamp time.Time       `json:"timestamp"`
	Payload   model.Product   `json:"payload"`
}
