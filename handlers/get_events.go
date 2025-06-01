package handlers

import (
	"events-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, models.EventsTests)
}
