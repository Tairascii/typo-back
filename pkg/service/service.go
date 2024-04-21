package service

import "typo_back/pkg/repository"

type Auth interface {
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
