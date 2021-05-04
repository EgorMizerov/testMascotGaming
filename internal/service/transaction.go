package service

import (
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository"
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
