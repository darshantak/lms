package handlers

import (
	"fmt"
	"lms/database"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadContent(c *gin.Context) {
	title := c.PostForm("title")
	fileType := c.PostForm("file_type")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", os.ModePerm)
	}

	filePath := filepath.Join("./uploads", file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	sqlStatement := fmt.Sprintf("INSERT INTO content (title, content_type, content_url, created_at, expires_at) VALUES ('%s', '%s', '%s','%s','%s')", title, fileType, filePath, time.Now().Format("2006-01-02 15:04:05"), time.Now().Add(100*time.Hour).Format("2006-01-02 15:04:05"))

	fmt.Println(sqlStatement)
	_, err = database.DB.Queryx(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert in the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "content uploaded"})
}
