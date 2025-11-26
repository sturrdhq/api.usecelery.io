package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"localhost:3000", "usecelery.io"})

	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":            true,
			"eaten veggies": true,
		})
	})

	router.Run()
}
