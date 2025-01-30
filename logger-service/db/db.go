package db

import (
	"context"
	"fmt"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/common"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	LogCollection *mongo.Collection
)

func ConnectDB() error {
	dbConfig := common.Config.DbConfig

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_ = client.Ping(ctx, readpref.Primary())

	LogCollection = client.Database(dbConfig.DbName).Collection(dbConfig.DbCollection)

	return nil
}
