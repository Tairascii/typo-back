package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"typo_back"
)

type UserDAO struct {
	c *mongo.Collection
}

func NewUserDAO(client *mongo.Client, dbName string, collection string) *UserDAO {
	return &UserDAO{c: client.Database(dbName).Collection(collection)}
}

func (dao *UserDAO) CreateUser(ctx context.Context, user typo_back.User) (int, error) {
	result, err := dao.c.InsertOne(ctx, user)

	if err != nil {
		log.Fatalf("something went wrong while inserting %s", err.Error())
	}
	log.Println(result)
	return 0, nil
}
