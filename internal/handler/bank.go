package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitBankRoutes(api *gin.RouterGroup) {
	banks := api.Group("/banks", h.userIdentify)
	{
		banks.POST("/", h.CreateBank)
		banks.GET("/", h.GetAllBanks)
	}
}

func (h *Handler) CreateBank(ctx *gin.Context) {
	id, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	bankId, err := h.service.Bank.CreateBank(id, "USD")
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bank_id": bankId,
	})
}

func (h *Handler) GetAllBanks(ctx *gin.Context) {
	_, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	banks, err := h.service.Bank.GetAllBanks()
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"banks": banks,
	})
}
