package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeAccess (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Access changed successfully"})
}

