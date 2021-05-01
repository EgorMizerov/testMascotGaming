package domain

type StartSessionResponse struct {
	Jsonrpc string
	Id      string
	Result  StartSessionResult
}

type StartSessionResult struct {
	SessionId  string
	SessionUrl string
}
