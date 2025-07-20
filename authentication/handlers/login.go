package handlers

import (
	"events-api/authentication/models"
	"events-api/authentication/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlers) LoginUser(context *gin.Context) {
	var user models.LoginUserDTO

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Provide username"})
		return
	}

	userDAO, err := h.UsersDB.GetUser(user.Username)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "User is not registered"})
		return
	}

	checkPasswordError := utils.CheckPassword(userDAO.Password, user.Password)

	if checkPasswordError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	accessToken, accessTokenGenerationError := utils.GenerateToken(uint(userDAO.ID))

	if accessTokenGenerationError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	refreshToken, refreshTokenGenerationError := utils.GenerateRefreshToken(32)

	if refreshTokenGenerationError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	saveRefreshTokenError := h.RefreshTokensDB.SaveRefreshToken(userDAO.ID, refreshToken)

	if saveRefreshTokenError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
