package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *EventHandlers) DeleteEvent(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID must be provided"})
	}

	err := h.DB.DeleteEvent(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
