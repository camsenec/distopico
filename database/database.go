package database

import (
	"context"
	"log"

	"github.com/distopico/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	session  *mongo.Client
	messages *mongo.Collection
}

func connectDB() DB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongo://example"))

	conf := config.ConfigModel{
		Uri:         "<URI>",
		Db:          "messages",
		TokenSecret: "secret",
		TokenExp:    "1h",
		ServeUri:    ":4444",
	}

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return DB{
		session:  client,
		messages: client.Database(conf.Db).Collection("messages"),
	}

}

func (db DB) closeDB() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
