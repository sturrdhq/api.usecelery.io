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
		Email string `json:"email" form:"email" binding:"required,email"`
	}

	var params subscribeParams
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "bad request"})
		return
	}

	err := h.service.Subscribe(params.Email)
	if err != nil {
		slog.Error("failed to register subscription", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to register subscription")
	}

	slog.Info(
		"Congratulations! You're in the waitlist",
	)

	c.JSON(http.StatusOK, "Congratulations! You're in the waitlist")
}
