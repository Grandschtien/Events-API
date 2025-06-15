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

	tx, err := h.DB.DB.Begin()

	if err != nil {
		log.Printf("Error on transaction creation %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err := h.DB.SaveUser(tx, user.Username, string(hash))

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			context.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}
		log.Printf("Error during saving of user: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	token, err := utils.GenerateToken(uint(id))

	if err != nil {
		log.Printf("Error while generating token %v", err)
		context.JSON(500, gin.H{"error": "internal error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"access_token": token})
}
