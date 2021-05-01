package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func ConnectionToPostgres(cfg Config) *sqlx.DB {
	query := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

	db, err := sqlx.Connect("pgx", query)
	if err != nil {
		log.Fatalf("database connection error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("ping error: %s", err)
	}

	return db
}
