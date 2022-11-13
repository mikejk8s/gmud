package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecureEndpoint(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is a secure endpoint"})
}
