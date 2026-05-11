package routes

import (
    "github.com/binojmadhuranga/BrainGrid-go/controllers"
    "github.com/binojmadhuranga/BrainGrid-go/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

    api := router.Group("/api")

    auth := api.Group("/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }

    protected := api.Group("/user")
    protected.Use(middleware.JWTAuthMiddleware())
    {
        protected.GET("/profile", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "Protected route",
            })
        })
    }
}