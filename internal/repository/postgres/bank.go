package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testMascotGaming/internal/domain"
)

type BankPostgres struct {
	db *sqlx.DB
}

func NewBankPostgres(db *sqlx.DB) *BankPostgres {
	return &BankPostgres{db: db}
}

func (r *BankPostgres) CreateBank(userId, bankId, currency string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO banks (id, user_id, currency) VALUES ($1, $2, $3)")
	_, err = tx.Exec(query, bankId, userId, currency)
	if err != nil {
		tx.Rollback()
		return err
	}

	id := uuid.New()
	query = fmt.Sprintf("INSERT INTO banks_users (id, user_id, bank_id) VALUES ($1, $2, $3)")
	_, err = tx.Exec(query, id.String(), userId, bankId)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}

func (r *BankPostgres) GetAllBanks() ([]domain.Bank, error) {
	var banks []domain.Bank

	query := fmt.Sprintf("SELECT * FROM banks")
	err := r.db.Select(&banks, query)

	return banks, err
}
