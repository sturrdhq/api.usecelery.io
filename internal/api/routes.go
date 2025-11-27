package api

import (
	"github.com/sturrdhq/celery-server/internal/api/handlers"
	"github.com/sturrdhq/celery-server/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer(p int, dbc *database.DBClient) {
	s := NewServer(p, dbc)

	defaultRouter := s.routerEngine
	defaultRouter.Use(CorsMiddleware())

	defaultRouter.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	defaultRouter.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "<h1>Welcome to Celery</h1>")
	})

	apiV1 := defaultRouter.Group("/api/v1")

	waitlistHandler := handlers.NewWaitListHandler(s.db)

	apiV1.PUT("/waitlist/subscribe", waitlistHandler.Subscribe)

	s.Start(p)
}
