package handlers

import (
	"database/sql"
	"lms/database"
	"lms/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func GetTraining(c *gin.Context) {
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

func CreateTraining(c *gin.Context) {
	var training models.Training
	userEmail := c.GetString("email")

	err := c.ShouldBindJSON(&training)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in training/CreateTraining/ShouldBindJSON"})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}
	sqlStatement := "INSERT INTO trainings (title, created_by, scoped_teams, scoped_content, expires_at) VALUES ($1, $2, $3, $4, $5);"

	_, err = database.DB.Exec(sqlStatement, training.Title, userEmail, pq.Array(training.ScopedTeams), pq.Array(training.ScopedContent), time.Now().Add(24*time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in training/CreateTraining/Exec"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Training successfully created"})
}
