package database

import (
	"context"
	"log"

	"github.com/distopico/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	session  *mongo.Client
	messages *mongo.Collection
}

func ConnectDB() MongoDB {
	conf := config.ConfigModel{
		Uri: "mongodb://172.17.0.2:27017",
		Db:  "messages",
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return MongoDB{
		session:  client,
		messages: client.Database(conf.Db).Collection("messages"),
	}

}

func (db MongoDB) CloseDB() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
