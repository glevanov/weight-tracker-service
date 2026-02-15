package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

const WeightsCollection = "weight"

func Connect(uri, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}

	if err := c.Ping(ctx, nil); err != nil {
		return err
	}

	client = c
	db = c.Database(dbName)

	return nil
}

func Disconnect(ctx context.Context) error {
	if client != nil {
		return client.Disconnect(ctx)
	}
	return nil
}

func GetWeightsCollection() *mongo.Collection {
	return db.Collection(WeightsCollection)
}
