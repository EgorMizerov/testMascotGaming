package service

import (
	"github.com/EgorMizerov/testMascotGaming/internal/client"
	"github.com/EgorMizerov/testMascotGaming/internal/repository"
	"github.com/spf13/viper"
)

type GameService struct {
	repo   repository.Bank
	client *client.Client
}

func NewGameService(repo repository.Bank, client2 *client.Client) *GameService {
	return &GameService{repo: repo, client: client2}
}

func (s *GameService) StartDemoGame(gameId string) (string, string, error) {
	bankId := viper.GetString("bankId")

	return s.client.StartDemoSession(bankId, gameId)
}

func (s *GameService) StartGame(gameId, userId string) (string, string, error) {
	return s.client.StartSession(userId, gameId)
}
