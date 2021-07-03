package database

import (
	"context"
	"time"

	"github.com/distopico/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db DB) GetAllMessages() (interface{}, error) {
	var results []model.MessageModel
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := db.messages.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var document model.MessageModel
		err = cur.Decode(&document)
		if err != nil {
			return nil, err
		}
		results = append(results, document)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)
	return results, nil
}

func (db DB) GetMessageById(_id string) (interface{}, error) {
	var result model.MessageModel
	var err error

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	q := bson.M{"id": id}
	err = db.messages.FindOne(ctx, q).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (db DB) createMessage(description string, topic string) (interface{}, error) {
	var document model.MessageModel
	var err error

	document.ID = primitive.NewObjectID()
	document.Topic = topic
	document.Description = description
	document.CreatedDate = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = db.messages.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return document, nil

}
