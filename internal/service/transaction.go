package service

import (
	"github.com/EgorMizerov/testMascotGaming/internal/domain"
	"github.com/EgorMizerov/testMascotGaming/internal/repository"
)

type TransactionService struct {
	repo    repository.Transaction
	balance repository.Balance
}

func NewTransactionService(repo repository.Transaction, balance repository.Balance) *TransactionService {
	return &TransactionService{repo: repo, balance: balance}
}

func (s *TransactionService) CreateTransaction(userId, ref string, withdraw, deposit int) (string, error) {
	return s.repo.CreateTransaction(userId, ref, withdraw, deposit)
}

func (s *TransactionService) GetTransactionByRef(ref string) (domain.Transaction, error) {
	return s.repo.GetTransactionByRef(ref)
}

func (s *TransactionService) Rollback(ref string) error {
	transaction, err := s.repo.GetTransactionByRef(ref)
	if err != nil {
		return err
	}

	if transaction.Withdraw != 0 {
		_, err = s.balance.Deposit(transaction.UserId, float64(transaction.Withdraw)/100)
		if err != nil {
			return err
		}
	}

	if transaction.Deposit != 0 {
		_, err = s.balance.Withdraw(transaction.UserId, float64(transaction.Deposit)/100)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *TransactionService) WithdrawAndDeposit(userId, transactionRef string, deposit, withdraw int) (float64, string, error) {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return 0, "", err
	}

	transactionId, err := s.repo.CreateTransactionDuringTransaction(tx, userId, transactionRef, withdraw, deposit)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}

	var balance float64

	if deposit != 0 {
		balance, err = s.balance.DepositDuringTransaction(tx, userId, float64(deposit)/100)
		if err != nil {
			tx.Rollback()
			return 0, "", err
		}
	}

	if withdraw != 0 {
		balance, err = s.balance.WithdrawDuringTransaction(tx, userId, float64(withdraw)/100)
		if err != nil {
			tx.Rollback()
			return 0, "", err
		}
	}

	tx.Commit()

	return balance, transactionId, err
}
