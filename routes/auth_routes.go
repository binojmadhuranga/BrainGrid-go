package routes

import (
	"github.com/binojmadhuranga/BrainGrid-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(api *gin.RouterGroup) {

	auth := api.Group("/auth")

	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}