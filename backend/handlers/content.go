package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadContent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "content uploaded"})
}

func GetContents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"contents": "list of contents"})
}
