package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"typo_back"
)

type UserDAO struct {
	c *mongo.Collection
}

func NewUserDAO(client *mongo.Client, dbName string, collection string) *UserDAO {
	return &UserDAO{c: client.Database(dbName).Collection(collection)}
}

func (dao *UserDAO) CreateUser(ctx context.Context, user typo_back.User) (primitive.ObjectID, error) {
	result, err := dao.c.InsertOne(ctx, user)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	return insertedID, nil
}
