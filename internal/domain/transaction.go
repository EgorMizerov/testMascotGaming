package domain

type Transaction struct {
	Id             string
	TransactionRef string
	UserId         string `db:"user_id"`
	Withdraw       int
	Deposit        int
}
