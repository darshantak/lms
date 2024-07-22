package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitAssessment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "assessment submitted"})
}

func GetAssessments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"assessments": "list of assessments"})
}
