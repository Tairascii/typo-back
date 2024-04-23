package repository

import "go.mongodb.org/mongo-driver/mongo"

type Auth interface {
}

type Repository struct {
	Auth
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{}
}
