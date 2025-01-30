package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()

	return
}

type Order struct {
	Base
	ProductId  uuid.UUID `json:"product_id"`
	Quantity   uint      `json:"quantity"`
	TotalPrice uint      `json:"total_price"`
	Status     string    `json:"status"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.Status = "unpaid"

	return o.Base.BeforeCreate(tx)
}

func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	return o.Base.BeforeUpdate(tx)
}
