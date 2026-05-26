package routes

import (
	"github.com/binojmadhuranga/BrainGrid-go/controllers"
	"github.com/binojmadhuranga/BrainGrid-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	// =========================
	// AUTH ROUTES
	// =========================

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// =========================
	// USER ROUTES
	// =========================

	user := api.Group("/user")

	user.Use(middleware.JWTAuthMiddleware())

	{
		user.GET("/profile", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "Protected user route",
			})
		})
	}

	// =========================
	// ADMIN ROUTES
	// =========================

	admin := api.Group("/admin")

	admin.Use(middleware.JWTAuthMiddleware())
	admin.Use(middleware.AdminMiddleware())

	{
		// Get all users
		admin.GET(
			"/users",
			controllers.GetUsers,
		)

		// Update user role
		admin.PUT(
			"/users/:id/role",
			controllers.UpdateUserRole,
		)

		// Update user status
		admin.PUT(
			"/users/:id/status",
			controllers.UpdateUserStatus,
		)

		// Delete user
		admin.DELETE(
			"/users/:id",
			controllers.DeleteUser,
		)
	}
}