package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonTokenOKResponse(context *gin.Context, userID int, accessToken string, refreshToken string) {
	context.JSON(http.StatusCreated, gin.H{
		"user_id":       userID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func CommonInternalErrorResponse(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}
