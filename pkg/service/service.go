package service

import (
	"context"
	"typo_back"
	"typo_back/pkg/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, user typo_back.User) (int, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
