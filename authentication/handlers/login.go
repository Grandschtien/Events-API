package handlers

import (
	"events-api/authentication/models"
	"events-api/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlers) LoginUser(context *gin.Context) {
	var user models.LoginUserDTO

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Provide username"})
		return
	}

	userDAO, err := h.DB.GetUser(user.Username)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "User is not registered"})
		return
	}

	checkPasswordError := utils.CheckPassword(userDAO.Password, user.Password)

	if checkPasswordError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, tokenGenerationError := utils.GenerateToken(uint(userDAO.ID))

	if tokenGenerationError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
