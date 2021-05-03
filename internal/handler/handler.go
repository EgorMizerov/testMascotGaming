package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"testMascotGaming/internal/auth"
	client2 "testMascotGaming/internal/client"
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

	router.GET("/ping", h.corsMiddleware, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	api := router.Group("/api", h.corsMiddleware)
	{
		h.InitAuthRoutes(api)
		h.InitUserRoutes(api)
		h.InitGamesRoutes(api)
		h.InitBankRoutes(api)
	}

	router.POST("/", func(ctx *gin.Context) {
		resp, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(resp))
	})

	return router
}
