package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongo(mongoURI string) *mongo.Client {
	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if mongoURI == "" {
		log.Fatal("MONGO_URI is empty")
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("✅ Connected to MongoDB")
	return client
}

func GetCollection(database string, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}