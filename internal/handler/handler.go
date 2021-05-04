package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"testMascotGaming/internal/auth"
	client2 "testMascotGaming/internal/client"
	domain2 "testMascotGaming/internal/client/domain"
	"testMascotGaming/internal/domain"
	"testMascotGaming/internal/service"
)

type Handler struct {
	service *service.Service
	log     *zap.Logger
	manager auth.TokenManager
	client  *client2.Client
}

func NewHandler(service *service.Service, log *zap.Logger, manager auth.TokenManager, client *client2.Client) *Handler {
	return &Handler{
		service: service,
		log:     log,
		manager: manager,
		client:  client,
	}
}

func (h *Handler) GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(h.optionMiddleware)
	router.Use(h.corsMiddleware)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	api := router.Group("/api")
	{
		h.InitAuthRoutes(api)
		h.InitUserRoutes(api)
		h.InitGamesRoutes(api)
		//h.InitBankRoutes(api)
	}

	router.POST("/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Accept", "application/json")

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			errorMessage(ctx, err, http.StatusBadRequest, err.Error())
			return
		}

		var JSONRPCReq domain.JSONRPCRequest
		err = json.Unmarshal(body, &JSONRPCReq)
		if err != nil {
			errorMessage(ctx, err, http.StatusBadRequest, err.Error())
			return
		}

		switch JSONRPCReq.Method {
		case "getBalance":
			h.getBalance(ctx, body)
		}
	})

	return router
}

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
