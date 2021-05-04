package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"testMascotGaming/internal/auth"
	client2 "testMascotGaming/internal/client"
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
	}

	router.POST("/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Accept", "application/json")

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			errorMessage(ctx, err, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(body)

		var JSONRPCReq domain.JSONRPCRequest
		err = json.Unmarshal(body, &JSONRPCReq)
		if err != nil {
			errorMessage(ctx, err, http.StatusBadRequest, err.Error())
			return
		}

		switch JSONRPCReq.Method {
		case "getBalance":
			h.getBalance(ctx, body)
		case "withdrawAndDeposit":
			h.withdrawAndDeposit(ctx, body)
		case "rollbackTransaction":
			h.rollbackTransaction(ctx, body)
		}
	})

	return router
}
