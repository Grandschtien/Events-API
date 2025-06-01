package handlers

import (
	"events-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	events := models.EventsTests // here we should retrieve data from DB
	id := c.Param("id")

	for _, event := range events {
		if event.ID == id {
			c.JSON(http.StatusOK, event)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
}
