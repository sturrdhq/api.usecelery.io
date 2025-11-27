package api

import (

	"strings"
	"time" // Added as per the provided Code Edit, though not explicitly used in the final snippet.

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://localhost:3000" {
				return true
			}
			if origin == "https://usecelery.io" || origin == "https://www.usecelery.io" {
				return true
			}
			// Allow Vercel previews
			if strings.HasSuffix(origin, ".vercel.app") {
				return true
			}
			// Allow ngrok tunnels
			if strings.HasSuffix(origin, ".ngrok-free.app") {
				return true
			}
			return false
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
