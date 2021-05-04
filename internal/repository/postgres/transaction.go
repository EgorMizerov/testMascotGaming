package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) CreateTransaction(userId, ref string, withdraw, deposit int) (string, error) {
	var id = generateId()

	query := fmt.Sprintf("INSERT INTO transactions (id, transactionRef, user_id, withdraw, deposit) VALUES ($1, $2, $3, $4, $5)")
	_, err := r.db.Exec(query, id, ref, userId, withdraw, deposit)

	return id, err
}

func generateId() string {
	id := uuid.New().String()

	args := strings.Split(id, "-")
	return "2:" + strings.Join(args[0:3], "")
}
