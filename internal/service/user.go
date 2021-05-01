package service

import (
	"go.uber.org/zap"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository"
)

type UserService struct {
	repo repository.User
	log  *zap.Logger
}

func NewUserService(repo repository.User, log *zap.Logger) *UserService {
	return &UserService{repo: repo, log: log}
}

func (s *UserService) GetUserByID(id string) (domain.User, error) {
	return s.repo.GetUserByID(id)
}
