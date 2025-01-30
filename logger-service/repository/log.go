package repository

import (
	"context"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const dbTimeout = 5 * time.Second

type logRepository struct {
	logCollection *mongo.Collection
}

type LogRepository interface {
	Save(log model.Log) (*model.Log, error)
	GetAll() ([]model.Log, error)
}

func NewLogRepository(logCollection *mongo.Collection) LogRepository {
	return &logRepository{
		logCollection: logCollection,
	}
}

func (l *logRepository) Save(log model.Log) (*model.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := l.logCollection.InsertOne(ctx, log)
	if err != nil {
		return nil, err
	}

	var savedUser model.Log

	filter := bson.M{
		"_id": result.InsertedID,
	}

	err = l.logCollection.FindOne(ctx, filter).Decode(&savedUser)
	if err != nil {
		return nil, err
	}

	return &savedUser, err
}

func (l *logRepository) GetAll() ([]model.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	cursor, err := l.logCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []model.Log

	for cursor.Next(ctx) {
		var log model.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}

		log.CreatedAt = log.CreatedAt.Local()
		log.UpdatedAt = log.UpdatedAt.Local()
		logs = append(logs, log)
	}

	return logs, nil
}
