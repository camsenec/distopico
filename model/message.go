package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Topic       string             `bson:"topic,required" json:"topic,required" form:"topic"`
	Description string             `bson:"description,required" json:"description,required" form:"description"`
	CreatedDate time.Time          `bson:"createdDate,required" json:"createdDate,required"`
}
