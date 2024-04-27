package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"typo_back"
	"typo_back/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user typo_back.User) (primitive.ObjectID, error) {
	return s.repo.CreateUser(ctx, user)
}
