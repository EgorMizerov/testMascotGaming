package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository/postgres"
)

type User interface {
	CreateUser(user domain.User) (string, error)
	GetUserByID(id string) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
	UpdateRefreshToken(id, token string) error
}

type Balance interface {
	Withdraw(id string, amount float64) (float64, error)
	Deposit(id string, amount float64) (float64, error)
}

type Bank interface {
	CreateBank(userId, bankId, currency string) error
	GetAllBanks() ([]domain.Bank, error)
}

type Repository struct {
	User
	Balance
	Bank
}

func NewRepository(db *sqlx.DB, log *zap.Logger) *Repository {
	return &Repository{
		User:    postgres.NewUserPostgres(db, log),
		Balance: postgres.NewBalancePostgres(db),
		Bank:    postgres.NewBankPostgres(db),
	}
}
