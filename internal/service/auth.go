package service

import (
	"errors"
	"go.uber.org/zap"
	"testMascotGaming/internal/auth"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository"
	"time"
)

type AuthService struct {
	repo    repository.User
	log     *zap.Logger
	manager auth.TokenManager
}

func NewAuthService(repo repository.User, log *zap.Logger, manager auth.TokenManager) *AuthService {
	return &AuthService{repo: repo, log: log, manager: manager}
}

func (s *AuthService) SignUp(password, username string) error {
	var user domain.User

	user.Username = username
	user.Password = HashPassword(password)

	_, err := s.repo.CreateUser(user)

	return err
}

func (s *AuthService) SignIn(password, username string) (string, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", "", errors.New("user is not found")
		}
		return "", "", err
	}

	if user.Password != HashPassword(password) {
		return "", "", errors.New("wrong password")
	}

	accessToken, err := s.manager.GenerateAccessToken(user.ID, false, time.Hour*2)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.manager.GenerateRefreshToken(user.ID, time.Hour*480)
	if err != nil {
		return "", "", err
	}

	err = s.repo.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func (s *AuthService) RefreshToken(id, token string) (string, string, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return "", "", err
	}

	// TODO: проверить валидность токена

	if user.RefreshToken.String != token {
		return "", "", errors.New("token invalid")
	}

	accessToken, err := s.manager.GenerateAccessToken(user.ID, false, time.Hour*2)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.manager.GenerateRefreshToken(user.ID, time.Hour*480)
	if err != nil {
		return "", "", err
	}

	err = s.repo.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}
