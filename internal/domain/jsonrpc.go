package domain

type JSONRPCRequest struct {
	Jsonrpc string `binding:"required"`
	Method  string `binding:"required"`
	Id      int
}
