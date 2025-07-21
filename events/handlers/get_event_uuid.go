package handlers

import (
	"database/sql"
	"net/http"

	coreUtils "events-api/core/utils"

	"github.com/gin-gonic/gin"
)

func (h *EventHandlers) GetEvent(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID must be provided"})
		return
	}

	event, err := h.DB.GetEvent(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}

		coreUtils.CommonInternalErrorResponse(c)
		return
	}

	c.JSON(http.StatusOK, event)
}
