package repository

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/model"
)

type esProductRepository struct {
	client    *elasticsearch.TypedClient
	indexName string
}

type ProductRepository interface {
	GetAllProducts() ([]model.Product, error)
	GetProductById(id string) (*model.Product, error)
}

func NewProductRepository(client *elasticsearch.TypedClient) ProductRepository {
	return &esProductRepository{
		client:    client,
		indexName: common.Config.ESConfig.IndexName,
	}
}

func (e *esProductRepository) GetAllProducts() ([]model.Product, error) {
	res, err := e.client.Search().Index(e.indexName).Size(10).From(0).Do(context.Background())
	if err != nil {
		return nil, err
	}

	var products []model.Product
	data := res.Hits.Hits

	for _, hit := range data {
		var product model.Product

		err := json.Unmarshal(hit.Source_, &product)
		if err != nil {
			return nil, err
		}

		product.ID = *hit.Id_

		products = append(products, product)
	}

	return products, nil
}

func (e *esProductRepository) GetProductById(id string) (*model.Product, error) {
	res, err := e.client.Get("products", id).Do(context.Background())
	if err != nil {
		return nil, err
	}

	var product model.Product

	err = json.Unmarshal(res.Source_, &product)
	if err != nil {
		return nil, err
	}
	product.ID = res.Id_

	return &product, nil
}
