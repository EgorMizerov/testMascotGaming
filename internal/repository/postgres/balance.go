package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) Withdraw(id string, amount float64) (float64, error) {
	query := fmt.Sprintf("UPDATE users SET balance=balance-$1 WHERE id=$2 RETURNING balance")

	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return 0, err
	}

	var balance float64

	query = fmt.Sprintf("SELECT balance FROM users WHERE id=$1")
	err = r.db.Get(&balance, query, id)

	return balance, err
}

func (r *BalancePostgres) Deposit(id string, amount float64) (float64, error) {
	query := fmt.Sprintf("UPDATE users SET balance=balance+$1 WHERE id=$2 RETURNING balance")

	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return 0, err
	}

	var balance float64

	query = fmt.Sprintf("SELECT balance FROM users WHERE id=$1")
	err = r.db.Get(&balance, query, id)

	return balance, err
}

func (r *BalancePostgres) GetBalance(id string) (float64, error) {
	var balance float64
	query := fmt.Sprintf("SELECT balance FROM users WHERE id=$1")

	err := r.db.Get(&balance, query, id)
	return balance, err
}
