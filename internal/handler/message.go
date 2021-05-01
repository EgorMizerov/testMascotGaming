package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"testMascotGaming/pkg/logger"
)

var log, _ = logger.NewLogger("debug")

func errorMessage(ctx *gin.Context, err error, status int, msg string) {
	log.Debug(msg, zap.Error(err))
	ctx.JSON(status, gin.H{
		"error": msg,
	})
}
