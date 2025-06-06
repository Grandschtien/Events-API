package handlers

import (
	"events-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handlers) AddEvent(c *gin.Context) {
	var event *models.EventAdd

	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unsupportable entity"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error with database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": internalEvent.UUID})
}
