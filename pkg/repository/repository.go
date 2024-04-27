package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"typo_back"
)

type Auth interface {
	CreateUser(ctx context.Context, user typo_back.User) (primitive.ObjectID, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{Auth: NewUserDAO(db, viper.GetString("db.name"), "user")}
}
