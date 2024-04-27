package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"typo_back"
	"typo_back/pkg/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, user typo_back.User) (primitive.ObjectID, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
