package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		setCorsHeaders(c, false)
	}
}

func setCorsHeaders(c *gin.Context, isOptions bool) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if isOptions {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
