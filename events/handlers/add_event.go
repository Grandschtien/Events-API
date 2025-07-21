package handlers

import (
	"events-api/events/models"
	"log"
	"net/http"

	coreUtils "events-api/core/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *EventHandlers) AddEvent(c *gin.Context) {
	var event *models.EventAdd

	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unsupportable entity"})
		return
	}

	if len(event.Title) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Title required"})
		return
	}

	internalEvent := models.Event{
		UUID:        uuid.New().String(),
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
	}

	_, err := h.DB.SaveEvent(internalEvent)

	if err != nil {
		log.Printf("Database error: %v", err)
		coreUtils.CommonInternalErrorResponse(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": internalEvent.UUID})
}
