package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SecureEndpoint(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is a secure endpoint"})
}
