package database

import (
	"context"
	"time"

	"weight-tracker-service/internal/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

const (
	WeightsCollection = "weight"
	UsersCollection   = "users"
)

func Connect(uri, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(clientOpts)
	if err != nil {
		logger.Error("database connection failed", "error", err)
		return err
	}

	if err := c.Ping(ctx, nil); err != nil {
		logger.Error("database ping failed", "error", err)
		return err
	}

	client = c
	db = c.Database(dbName)

	logger.Info("database connected", "database", dbName)
	return nil
}

func Disconnect(ctx context.Context) error {
	if client != nil {
		err := client.Disconnect(ctx)
		if err != nil {
			logger.Error("database disconnect failed", "error", err)
			return err
		}
		logger.Info("database disconnected")
		return nil
	}
	return nil
}

func GetWeightsCollection() *mongo.Collection {
	return db.Collection(WeightsCollection)
}

func GetUsersCollection() *mongo.Collection {
	return db.Collection(UsersCollection)
}
