package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	err := h.DB.DeleteEvent(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
