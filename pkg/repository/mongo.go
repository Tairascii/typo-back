package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoDB(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("something went wrong while disconnecting %s", err.Error())
		}
	}()

	return client, nil
}
