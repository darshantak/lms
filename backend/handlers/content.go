package handlers

import (
	"database/sql"
	"fmt"
	"lms/database"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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

func GetContent(c *gin.Context) {
	userEmail := c.GetString("email")
	log.Printf("Retrieved email from context: %s", userEmail)

	var storedTrainingArray []string

	sqlStatement := "SELECT training_assigned FROM users WHERE email=$1"

	// err := database.DB.Select(&storedTrainingArray, sqlStatement, userEmail)
	err := database.DB.QueryRow(sqlStatement, userEmail).Scan(pq.Array(&storedTrainingArray))
	if err != nil {
		log.Printf("Database query error: %v", err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"path": "content/GetContent", "error": "No records found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"path": "content/GetContent", "error": "Failed to fetch rows from DB"})
		return
	}

	log.Printf("Retrieved training assigned: %v", storedTrainingArray)

	c.JSON(http.StatusOK, gin.H{"contents": storedTrainingArray})
}
