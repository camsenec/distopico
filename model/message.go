package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageModel struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ApplicationID uint64             `bson:"applicationId,omitempty" json:"applicationId,omitempty"`
	Description   string             `bson:"description" json:"description,omitempty"`
	Topic         string             `bson:"topic" json:"topic,omitempty"`
	CreatedData   time.Time          `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
}
