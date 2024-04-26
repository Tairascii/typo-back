package service

import (
	"context"
	"typo_back"
	"typo_back/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user typo_back.User) (int, error) {
	return s.repo.CreateUser(ctx, user)
}
