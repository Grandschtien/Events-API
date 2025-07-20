package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hander *AuthHandlers) RefreshToken(context *gin.Context) {
	var token string
	var userID int

	if err := context.BindJSON(&token); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Token is empty"})
		return
	}

	if err := context.BindJSON(&userID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User id is empty"})
		return
	}

}
