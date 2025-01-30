package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Base struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
	DeletedAt *time.Time    `bson:"deleted_at"`
}

type Log struct {
	Base    `bson:"inline"`
	Content string `bson:"content"`
}

func NewLog(content string) Log {
	return Log{
		Base: Base{
			ID:        bson.NewObjectIDFromTimestamp(time.Now()),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Content: content,
	}
}
