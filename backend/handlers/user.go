package handlers

import (
	"fmt"
	"lms/database"
	"lms/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	//implement user reg
	var user models.User
	role := "ADMIN"
	empType := "FULL"
	team := "Security"
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	sqlStatementForAuth := fmt.Sprintf("INSERT INTO user_auth (email,password_hash,role_type,registered_at) VALUES ('%s','%s','%s','%s')", user.Email, string(hashedPassword), role, time.Now().Format("2006-01-02 15:04:05"))

	result, err := database.DB.Exec(sqlStatementForAuth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert the values"})
		return
	}

	_, err = result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get authId"})
		return
	}

	sqlStatementForLms := fmt.Sprintf("INSERT INTO users (username, role, emp_type, team, training_assigned, training_completed, created_at) VALUES ('%s','%s','%s','%s','%s','%s','%s')", user.Username, role, empType, team, "{}", "{}", time.Now().Format("2006-01-02 15:04:05"))

	_, err = database.DB.Exec(sqlStatementForLms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert in the lms table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registration successfull"})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user logged in"})
}
