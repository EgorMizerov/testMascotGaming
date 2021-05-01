package domain

type Bank struct {
	Id       string
	Currency string
	UserID   string `db:"user_id"`
}
