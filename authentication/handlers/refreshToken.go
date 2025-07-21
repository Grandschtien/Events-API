package handlers

import (
	"events-api/authentication/models"
	"events-api/authentication/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (hander *AuthHandlers) RefreshToken(context *gin.Context) {
	var token models.Tokens

	if err := context.BindJSON(&token); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Token is empty"})
		return
	}

	err, tokens := hander.RefreshTokensDB.GetRefreshTokens(token.UserID)

	if err != nil {
		utils.CommonInternalErrorResponse(context)
		return
	}

	commonAuthError := func() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Required authorization"})
	}

	if len(tokens) == 0 {
		commonAuthError()
		return
	}

	latestToken := selectLatestToken(tokens)

	if token.RefreshToken != latestToken.TokenHash {
		commonAuthError()
		return
	}

	now := time.Now()

	if latestToken.ExpiresAt.Unix() < now.Unix() {
		hander.RefreshTokensDB.RevokeRefreshToken(token.UserID, true)
		commonAuthError()
		return
	}

	if latestToken.Revoked {
		commonAuthError()
		return
	}

	newAccessToken, tokenGenerationError := utils.GenerateToken(uint(token.UserID))

	if tokenGenerationError != nil {
		utils.CommonInternalErrorResponse(context)
		return
	}

	utils.CommonTokenOKResponse(context, token.UserID, newAccessToken, token.RefreshToken)
}

func selectLatestToken(tokens []models.RefreshTokenDAO) models.RefreshTokenDAO {
	latestToken := tokens[0]

	for i := 1; i < len(tokens); i++ {
		if tokens[i].IssuedAt.Unix() > latestToken.IssuedAt.Unix() {
			latestToken = tokens[i]
		}
	}

	return latestToken
}
