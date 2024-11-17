package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Database struct {
	Collection *mongo.Collection
}

func Connect(ctx context.Context) (*mongo.Client, *Database, error) {
	mongoDbUri := os.Getenv("MONGODB_URI")
	if mongoDbUri == "" {
		return nil, nil, fmt.Errorf("MONGODB_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoDbUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	fmt.Println("Connected to MONGODB Atlas")

	db := &Database{
		Collection: client.Database("golang_db").Collection("todos"),
	}

	return client, db, nil
}
