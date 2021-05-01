package service

import "testMascotGaming/internal/repository"

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) Withdraw(id string, amount float64) (float64, error) {
	return s.repo.Withdraw(id, amount)
}

func (s *BalanceService) Deposit(id string, amount float64) (float64, error) {
	return s.repo.Deposit(id, amount)
}
