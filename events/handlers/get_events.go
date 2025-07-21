package handlers

import (
	"net/http"

	coreUtils "events-api/core/utils"

	"github.com/gin-gonic/gin"
)

func (h *EventHandlers) GetEvents(c *gin.Context) {
	events, err := h.DB.GetAllEvents()

	if err != nil {
		coreUtils.CommonInternalErrorResponse(c)
		return
	}
	c.JSON(http.StatusOK, events)
}
