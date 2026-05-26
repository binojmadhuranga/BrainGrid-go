package routes

import (
	"github.com/binojmadhuranga/BrainGrid-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup) {

	user := api.Group("/user")

	user.Use(middleware.JWTAuthMiddleware())

	{
		user.GET("/profile", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "Protected user route",
			})
		})
	}
}