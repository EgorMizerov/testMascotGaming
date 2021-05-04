package service

import "testMascotGaming/internal/repository"

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userId, ref string, withdraw, deposit int) (string, error) {
	return s.repo.CreateTransaction(userId, ref, withdraw, deposit)
}
