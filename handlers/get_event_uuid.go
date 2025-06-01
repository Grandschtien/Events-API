package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.DB.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}
