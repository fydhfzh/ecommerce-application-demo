package repository

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/dto"
)

type esProductRepository struct {
	client    *elasticsearch.TypedClient
	indexName string
}

type ProductRepository interface {
	CreateProductIndex() error
	SaveProduct(product dto.Product) error
}

func NewESRepository(client *elasticsearch.TypedClient) ProductRepository {
	return &esProductRepository{
		client:    client,
		indexName: common.Config.ESConfig.IndexName,
	}
}

func (e *esProductRepository) CreateProductIndex() error {
	exists, err := e.client.Indices.Exists(e.indexName).Do(context.Background())
	if err != nil {
		return err
	}

	if exists {
		log.Printf("%s index already exists", e.indexName)
		return nil
	}

	res, err := e.client.Indices.Create(e.indexName).
		Request(
			&create.Request{
				Mappings: &types.TypeMapping{
					Properties: map[string]types.Property{
						"name":        types.NewKeywordProperty(),
						"description": types.NewTextProperty(),
						"stock":       types.NewIntegerNumberProperty(),
						"created_at":  types.NewDateProperty(),
						"updated_at":  types.NewDateProperty(),
					},
				},
			},
		).
		Do(context.Background())

	if err != nil {
		return err
	}

	log.Printf("[INFO] %v\n", res)

	return nil
}

func (e *esProductRepository) SaveProduct(product dto.Product) error {
	res, err := e.client.Index(e.indexName).Request(product).Do(context.Background())
	if err != nil {
		return err
	}

	log.Printf("[INFO] %v\n", res)
	return nil
}
