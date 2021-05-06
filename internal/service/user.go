package service

import (
	"github.com/EgorMizerov/testMascotGaming/internal/domain"
	"github.com/EgorMizerov/testMascotGaming/internal/repository"
	"go.uber.org/zap"
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
