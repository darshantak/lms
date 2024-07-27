package routes

import (
	"lms/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", handlers.LoginUser)
	r.POST("/api/register", handlers.RegisterUser)

	generalUser := r.Group("/api").Use(handlers.Authenticate())
	{
		generalUser.POST("/api/contents/upload", handlers.UploadContent)
		generalUser.GET("/api/contents", handlers.GetContent)
		generalUser.GET("/api/assessments", handlers.GetAssessments)
		generalUser.POST("/api/assessments", handlers.CreateAssessments)
		generalUser.POST("/api/assessments/submit", handlers.SubmitAssessment)
	}
	adminUser := r.Group("/api/admin").Use(handlers.Authenticate())
	{
		adminUser.POST("/api/admin/changeAccess", handlers.ChangeAccess)
	}
	return r
}
