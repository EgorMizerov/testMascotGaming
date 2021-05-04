package domain

type StartSessionResponse struct {
	Jsonrpc string
	Id      int
	Result  StartSessionResult
}

type StartSessionResult struct {
	SessionId  string `json:"SessionId"`
	SessionUrl string `json:"SessionUrl"`
}
