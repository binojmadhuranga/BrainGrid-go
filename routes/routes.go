package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	RegisterAuthRoutes(api)
	RegisterUserRoutes(api)
	RegisterAdminRoutes(api)
}