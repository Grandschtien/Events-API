package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetEvents(c *gin.Context) {
	events, err := h.DB.GetAllEvents()

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, events)
}
