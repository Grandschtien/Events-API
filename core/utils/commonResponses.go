package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonInternalErrorResponse(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}
