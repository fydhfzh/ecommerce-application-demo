package dto

import (
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/constant"
)

type ProductEvent struct {
	EventType constant.Method `json:"event_type"`
	ProductId string          `json:"product_id"`
	Timestamp time.Time       `json:"timestamp"`
	Payload   Product         `json:"payload"`
}

type Product struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Stock       uint       `json:"stock"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
