package src

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	stickerscol *mongo.Collection
	chatcol     *mongo.Collection
)

func init() {
	log.Println("Setting up database...")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(Envars.DbUrl))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	stickerscol = client.Database("stickersbot").Collection("stickers")
	chatcol = client.Database("stickersbot").Collection("chat")
}
