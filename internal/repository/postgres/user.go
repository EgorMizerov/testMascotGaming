package postgres

import (
	"fmt"
	"github.com/EgorMizerov/testMascotGaming/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserPostgres struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewUserPostgres(db *sqlx.DB, log *zap.Logger) *UserPostgres {
	return &UserPostgres{db: db, log: log}
}

func (r *UserPostgres) CreateUser(user domain.User) (string, error) {
	userUUID := uuid.New()

	query := fmt.Sprintf("INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3)")
	_, err := r.db.Exec(query, userUUID.String(), user.Username, user.Password)

	return userUUID.String(), err
}

func (r *UserPostgres) GetUserByID(id string) (domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT * FROM users WHERE id=$1")
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserPostgres) UpdateRefreshToken(id, token string) error {
	query := fmt.Sprintf("UPDATE users SET refresh_token=$1 WHERE id=$2")
	_, err := r.db.Exec(query, token, id)

	return err
}

func (r *UserPostgres) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT * FROM users WHERE username=$1")
	err := r.db.Get(&user, query, username)

	return user, err
}
