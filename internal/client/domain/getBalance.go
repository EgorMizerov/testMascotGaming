package domain

type GetBalanceRequest struct {
	Jsonrpc string
	Method  string
	Id      string
	Params  GetBalanceRequestParams
}

type GetBalanceRequestParams struct {
	CallerId   int    `json:"callerId"`
	PlayerName string `json:"playerName"`
	Currency   string
	GameId     string `json:"gameId"`
	SessionId  string `json:"sessionId"`
}
