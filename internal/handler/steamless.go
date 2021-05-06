package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	domain2 "testMascotGaming/internal/client/domain"
)

func (h *Handler) getBalance(ctx *gin.Context, body []byte) {
	var getBalanceReq domain2.GetBalanceRequest

	err := json.Unmarshal(body, &getBalanceReq)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	balance, err := h.service.Balance.GetBalance(getBalanceReq.Params.PlayerName)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}
	fmtBalance := int(balance * 100)

	ctx.JSON(http.StatusOK, gin.H{
		"jsonrpc": "2.0",
		"id":      getBalanceReq.Id,
		"result": map[string]int{
			"balance": fmtBalance,
		},
	})
}

func (h *Handler) withdrawAndDeposit(ctx *gin.Context, body []byte) {
	var withdrawAndDepositReq withdrawAndDepositRequest

	err := json.Unmarshal(body, &withdrawAndDepositReq)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	var userId = withdrawAndDepositReq.Params.PlayerName
	var balance float64

	balance, transactionId, err := h.service.Transaction.WithdrawAndDeposit(
		userId,
		withdrawAndDepositReq.Params.TransactionRef,
		withdrawAndDepositReq.Params.Deposit,
		withdrawAndDepositReq.Params.Withdraw)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	fmtBalance := int(balance * 100)
	result := map[string]interface{}{
		"newBalance":    fmtBalance,
		"transactionId": transactionId,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"jsonrpc": "2.0",
		"id":      withdrawAndDepositReq.Id,
		"result":  result,
	})
}

type withdrawAndDepositRequest struct {
	Jsonrpc string
	Method  string
	Id      int
	Params  withdrawAndDepositParams
}

type withdrawAndDepositParams struct {
	PlayerName     string `json:"playerName"`
	Withdraw       int
	Deposit        int
	TransactionRef string `json:"transactionRef"`
}

func (h *Handler) rollbackTransaction(ctx *gin.Context, body []byte) {
	var rollbackTransactionReq rollbackTransactionRequest

	err := json.Unmarshal(body, &rollbackTransactionReq)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Rollback(rollbackTransactionReq.Params.TransactionRef)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"jsonrpc": "2.0",
		"id":      rollbackTransactionReq.Id,
		"result":  map[string]interface{}{},
	})
}

type rollbackTransactionRequest struct {
	Jsonrpc string
	Method  string
	Id      int
	Params  rollbackTransactionParams
}

type rollbackTransactionParams struct {
	PlayerName     string `json:"playerName"`
	TransactionRef string `json:"transactionRef"`
}
