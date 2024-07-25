package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context){
	//implement user reg
	
	c.JSON(http.StatusOK,gin.H{"message":"user registration successfull"})
}

func LoginUser(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"user logged in"})
}

