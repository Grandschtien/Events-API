package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *EventHandlers) GetEvents(c *gin.Context) {
	events, err := h.DB.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, events)
}
