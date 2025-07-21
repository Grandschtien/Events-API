package handlers

import (
	"events-api/authentication/models"
	"events-api/authentication/utils"
	coreUtils "events-api/core/utils"

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
		coreUtils.CommonInternalErrorResponse(context)
		return
	}

	refreshToken, refreshTokenGenerationError := utils.GenerateRefreshToken(32)

	if refreshTokenGenerationError != nil {
		coreUtils.CommonInternalErrorResponse(context)
		return
	}

	revokeTokensError := h.RefreshTokensDB.RevokeRefreshToken(userDAO.ID, true)

	if revokeTokensError != nil {
		coreUtils.CommonInternalErrorResponse(context)
		return
	}

	saveRefreshTokenError := h.RefreshTokensDB.SaveRefreshToken(userDAO.ID, refreshToken)

	if saveRefreshTokenError != nil {
		coreUtils.CommonInternalErrorResponse(context)
		return
	}

	utils.CommonTokenOKResponse(context, userDAO.ID, accessToken, refreshToken)
}
