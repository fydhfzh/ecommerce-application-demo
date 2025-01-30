package repository

import (
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/entity"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

type OrderRepository interface {
	CreateOrder(order entity.Order) (*entity.Order, error)
	GetOrderById(id int) (*entity.Order, error)
}

func NewRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (o *orderRepository) CreateOrder(newOrder entity.Order) (*entity.Order, error) {
	res := o.db.Create(&newOrder)
	if err := res.Error; err != nil {
		return nil, err
	}

	return &newOrder, nil
}

func (o *orderRepository) GetOrderById(id int) (*entity.Order, error) {
	var order entity.Order

	res := o.db.First(&order, id)
	if err := res.Error; err != nil {
		return nil, err
	}

	return &order, nil
}
