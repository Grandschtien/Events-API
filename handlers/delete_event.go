package handlers

import (
	"events-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	for i, val := range models.EventsTests {
		if val.ID == id {
			models.EventsTests = append(models.EventsTests[:i], models.EventsTests[i+1:]...)
			c.JSON(http.StatusOK, gin.H{})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
}
