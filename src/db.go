package src

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client  *mongo.Client
	chatcol *mongo.Collection
)

func init() {
	log.Println("Setting up database...")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(Envars.DbUrl))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	chatcol = client.Database("bot").Collection("chat")
}
