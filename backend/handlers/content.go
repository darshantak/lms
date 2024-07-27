package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadContent(c *gin.Context) {
	userEmail := c.GetString("email")
	fmt.Println(userEmail)

	c.JSON(http.StatusOK, gin.H{"message": "content uploaded"})
}

func GetContent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"contents": "list of contents"})
}
