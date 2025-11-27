package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sturrdhq/celery-server/internal/database"
	"github.com/sturrdhq/celery-server/internal/service/waitlist"
)

type WaitlistHandler struct {
	service *waitlist.Service
}

func NewWaitListHandler(db *database.DBClient) *WaitlistHandler {
	s := waitlist.NewWaitListService(db)
	h := &WaitlistHandler{s}
	return h
}

func (h *WaitlistHandler) Subscribe(c *gin.Context) {
	type subscribeParams struct {
		Email string `form:"email"`
	}

	var params subscribeParams
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "bad request"})
		return
	}

	err := h.service.Subscribe(params.Email)
	if err != nil {
		slog.Error("failed to register subscription", "err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Failed to register subscription"})
		return
	}

	slog.Info(
		"Congratulations! You're in the waitlist",
	)

	c.JSON(http.StatusOK, gin.H{"message": "Congratulations! You're in the waitlist", "error": nil})
}
