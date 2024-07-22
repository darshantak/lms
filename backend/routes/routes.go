package routes

import (
	"lms/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.POST("/content/upload", handlers.UploadContent)
	r.GET("/contents", handlers.GetAssessments)
	r.POST("/assessments/submit", handlers.SubmitAssessment)
	r.POST("/assessments", handlers.GetAssessments)

	return r
}
