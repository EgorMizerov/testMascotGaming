package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)

		authorized := auth.Group("", h.userIdentify)
		{
			authorized.POST("/refresh", h.Refresh)
		}
	}
}

type inputByCreateUser struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var input inputByCreateUser

	err := ctx.BindJSON(&input)
	if err != nil {
		errorMessage(ctx, err, http.StatusBadRequest, "invalid input body")
		return
	}

	at, rt, err := h.service.Auth.SignIn(input.Password, input.Username)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  at,
		"refresh_token": rt,
	})
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var input inputByCreateUser

	err := ctx.BindJSON(&input)
	if err != nil {
		errorMessage(ctx, err, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.service.Auth.SignUp(input.Password, input.Username)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user is created",
	})
}

type inputRefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (h *Handler) Refresh(ctx *gin.Context) {
	var input inputRefreshToken

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

	at, rt, err := h.service.Auth.RefreshToken(id, input.RefreshToken)
	if err != nil {
		errorMessage(ctx, err, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  at,
		"refresh_token": rt,
	})
}
