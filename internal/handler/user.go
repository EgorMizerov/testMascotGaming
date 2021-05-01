package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users", h.userIdentify)
	{
		users.GET("/", h.GetUserData)
		users.POST("/balance/withdraw", h.Withdraw)
		users.POST("/balance/deposit", h.Deposit)
	}
}

func (h *Handler) GetUserData(ctx *gin.Context) {
	id, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.User.GetUserByID(id)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"balance":  user.Balance,
	})
}

type inputSum struct {
	Amount float64
}

func (h *Handler) Withdraw(ctx *gin.Context) {
	var input inputSum

	id, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	err = ctx.BindJSON(&input)
	if err != nil {
		errorMessage(ctx, err, http.StatusBadRequest, "invalid input body")
		return
	}

	balance, err := h.service.Balance.Withdraw(id, input.Amount)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}

func (h *Handler) Deposit(ctx *gin.Context) {
	var input inputSum

	id, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	err = ctx.BindJSON(&input)
	if err != nil {
		errorMessage(ctx, err, http.StatusBadRequest, "invalid input body")
		return
	}

	balance, err := h.service.Balance.Deposit(id, input.Amount)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
