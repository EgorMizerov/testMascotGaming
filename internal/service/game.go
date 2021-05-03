package service

import (
	"testMascotGaming/internal/client"
	"testMascotGaming/internal/repository"
)

type GameService struct {
	repo   repository.Bank
	client *client.Client
}

func NewGameService(repo repository.Bank, client2 *client.Client) *GameService {
	return &GameService{repo: repo, client: client2}
}

func (s *GameService) StartDemoGame(gameId, userId string) (string, string, error) {
	// TODO: поменять на найти банк по id пользователя
	banks, _ := s.repo.GetAllBanks()
	return s.client.StartDemoSession(banks[0].Id, gameId)
}
