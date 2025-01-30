package db

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
)

var (
	Client *elasticsearch.TypedClient
)

func ConnectDB() error {
	esConfig := common.Config.ESConfig

	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%d", esConfig.Host, esConfig.Port),
		},
	}

	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return err
	}

	Client = es

	return err
}
