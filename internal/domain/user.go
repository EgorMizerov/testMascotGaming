package domain

import "database/sql"

type User struct {
	ID           string
	Username     string
	Password     string `db:"password_hash"`
	Balance      float64
	RefreshToken sql.NullString `db:"refresh_token"`
}
