package handlers

import (
	"errors"
	"events-api/authentication/models"
	"events-api/authentication/utils"
	"log"
	"net/http"

	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandlers) RegisterUser(context *gin.Context) {
	var user *models.RegisterUser

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unsupportable entity"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error of hashing password %v", err)
		context.JSON(500, gin.H{"error": "internal error"})
		return
	}

	tx, err := h.UsersDB.DB.Begin()

	if err != nil {
		log.Printf("Error on transaction creation %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	id, err := h.UsersDB.SaveUser(tx, user.Username, string(hash))

	if err != nil {
		_ = tx.Rollback()

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			context.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}
		log.Printf("Error during saving of user: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	tx.Commit()

	accessToken, accessTokenGenerationError := utils.GenerateToken(uint(id))

	if accessTokenGenerationError != nil {
		log.Printf("Error while generating access token %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	refreshToken, refreshTokenGenerationError := utils.GenerateRefreshToken(32)

	if refreshTokenGenerationError != nil {
		log.Printf("Error while generating refresh token %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	refreshTokenSavingError := h.RefreshTokensDB.SaveRefreshToken(id, refreshToken)

	if refreshTokenSavingError != nil {
		log.Printf("Error while saving refresh token %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Need authorization"})
		return
	}

	utils.CommonTokenOKResponse(context, id, accessToken, refreshToken)
}
