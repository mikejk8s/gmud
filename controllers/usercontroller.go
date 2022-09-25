package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/userdb"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := userdb.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userID": user.ID, "username": user.Username, "email": user.Email})
}

func GetUser(context *gin.Context) {
	var user models.User
	userID := context.Param("userID")
	record := userdb.Instance.First(&user, userID)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"userID": user.ID, "username": user.Username, "email": user.Email})
}