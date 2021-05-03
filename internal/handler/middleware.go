package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) parseAuthHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.manager.Parse(headerParts[1])
}

func (h *Handler) userIdentify(ctx *gin.Context) {
	id, err := h.parseAuthHeader(ctx)
	if err != nil {
		return
	}

	ctx.Set(userCtx, id)
}

func getUserID(c *gin.Context) (string, error) {
	idFromCtx, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("empty auth header")
	}

	id, ok := idFromCtx.(string)
	if !ok {
		return "", errors.New("userCtx is of invalid type")
	}

	return id, nil
}

func (h *Handler) corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
