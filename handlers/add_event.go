package handlers

import (
	"events-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEvent(c *gin.Context) {
	var event *models.EventDTO

	if err := c.BindJSON(&event); err != nil {
		// throw error that we
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unsupportable entity"})
		return
	}

	models.EventsTests = append(models.EventsTests, *event)
	c.JSON(http.StatusCreated, event)
}
