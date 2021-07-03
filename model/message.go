package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Topic       string             `bson:"topic" json:"topic,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	CreatedDate time.Time          `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
}
