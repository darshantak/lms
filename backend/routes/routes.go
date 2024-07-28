package routes

import (
	"lms/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	r.POST("/api/login", handlers.LoginUser)
	r.POST("/api/register", handlers.RegisterUser)

	generalUser := r.Group("/api").Use(handlers.Authenticate())
	{
		generalUser.POST("/contents/upload", handlers.UploadContent)
		generalUser.GET("/training", handlers.GetTraining)
		generalUser.POST("/training", handlers.CreateTraining)
		generalUser.GET("/assessments", handlers.GetAssessments)
		generalUser.POST("/assessments", handlers.CreateAssessments)
		generalUser.POST("/assessments/submit", handlers.SubmitAssessment)
	}
	adminUser := r.Group("/api/admin").Use(handlers.Authenticate())
	{
		adminUser.POST("/changeAccess", handlers.ChangeAccess)
	}
	return r
}
