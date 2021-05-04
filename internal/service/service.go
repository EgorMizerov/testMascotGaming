package service

import (
	"crypto/sha256"
	"fmt"
	"go.uber.org/zap"
	"testMascotGaming/internal/auth"
	"testMascotGaming/internal/client"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository"
)

type User interface {
	GetUserByID(id string) (domain.User, error)
}

type Auth interface {
	SignUp(password, username string) error
	SignIn(password, username string) (string, string, error)
	RefreshToken(id, token string) (string, string, error)
}

type Balance interface {
	GetBalance(id string) (float64, error)
	Withdraw(id string, amount float64) (float64, error)
	Deposit(id string, amount float64) (float64, error)
}

type Bank interface {
	CreateBank(userId, currency string) (string, error)
	GetAllBanks() ([]domain.Bank, error)
	SetPlayer(userId, username string) error
}

type Game interface {
	StartDemoGame(gameId string) (string, string, error)
	StartGame(gameId, playerId string) (string, string, error)
}

type Transaction interface {
	CreateTransaction(userId, ref string, withdraw, deposit int) (string, error)
	Rollback(ref string) error
}

type Service struct {
	User
	Auth
	Balance
	Bank
	Game
	Transaction
}

func NewService(repo *repository.Repository, log *zap.Logger, manager auth.TokenManager, client *client.Client) *Service {
	return &Service{
		Auth:        NewAuthService(repo.User, log, manager, client),
		User:        NewUserService(repo.User, log),
		Balance:     NewBalanceService(repo.Balance),
		Bank:        NewBankService(repo.Bank, repo.User, client),
		Game:        NewGameService(repo.Bank, client),
		Transaction: NewTransactionService(repo.Transaction, repo.Balance),
	}
}

const SALT = "qgc&^CWB7^GEVc7egdsvbucyVY&Q^WFGcvs`uydzhvgcsID^&Fgcisv7`bd jkhcb`wea8y"

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password + SALT))
	ph := h.Sum(nil)

	return fmt.Sprintf("%x", ph)
}
