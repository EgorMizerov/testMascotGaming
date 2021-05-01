package service

import (
	"testMascotGaming/internal/client"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/repository"
)

type BankService struct {
	banks  repository.Bank
	users  repository.User
	client *client.Client
}

func NewBankService(repo repository.Bank, users repository.User, client *client.Client) *BankService {
	return &BankService{banks: repo, users: users, client: client}
}

func (s *BankService) CreateBank(userId, currency string) (string, error) {
	user, err := s.users.GetUserByID(userId)
	if err != nil {
		return "", err
	}

	bankId, err := s.client.SetBankGroup(userId)
	if err != nil {
		return "", err
	}

	err = s.client.SetPlayer(userId, user.Username, bankId)
	if err != nil {
		return "", err
	}

	err = s.banks.CreateBank(userId, bankId, currency)

	return bankId, err
}

func (s *BankService) GetAllBanks() ([]domain.Bank, error) {
	return s.banks.GetAllBanks()
}
