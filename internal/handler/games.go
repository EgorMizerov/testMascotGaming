package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitGamesRoutes(api *gin.RouterGroup) {
	games := api.Group("/games", h.userIdentify)
	{
		games.GET("/", h.GetGameList)
		games.POST("/start", h.StartSession)
	}
}

func (h *Handler) GetGameList(ctx *gin.Context) {
	list, err := h.client.GetGameList()
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

type inputGame struct {
	GameID string `json:"game_id" binding:"required"`
}

func (h *Handler) StartSession(ctx *gin.Context) {
	id, err := getUserID(ctx)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	var input inputGame
	err = ctx.BindJSON(&input)
	if err != nil {
		errorMessage(ctx, err, http.StatusBadRequest, "invalid input body")
		return
	}

	gameId, gameUrl, err := h.client.StartSession(id, input.GameID)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"game_id":  gameId,
		"game_url": gameUrl,
	})
}
