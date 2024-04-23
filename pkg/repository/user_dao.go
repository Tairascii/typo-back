package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO struct {
	c *mongo.Collection
}

func NewUserDAO(ctx context.Context, client *mongo.Client, dbName string, collection string) *UserDAO {
	return &UserDAO{c: client.Database(dbName).Collection(collection)}
}
