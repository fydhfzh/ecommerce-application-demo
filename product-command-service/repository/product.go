package repository

import (
	"log"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Save(product model.Product) (*model.Product, error)
	GetAll() ([]model.Product, error)
	GetOne(id string) (*model.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) Save(product model.Product) (*model.Product, error) {
	if result := p.db.Save(&product); result.Error != nil {
		return nil, result.Error
	}

	log.Println(product)

	return &product, nil
}

func (p *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if result := p.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *productRepository) GetOne(id string) (*model.Product, error) {
	var product model.Product

	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if result := p.db.First(&product, "id = ?", uuidId); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
