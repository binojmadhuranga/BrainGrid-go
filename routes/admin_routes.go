package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/binojmadhuranga/BrainGrid-go/controllers"
	"github.com/binojmadhuranga/BrainGrid-go/middleware"
)

func AdminRoutes(r *gin.RouterGroup) {

	admin := r.Group("/admin")

	admin.Use(middleware.JWTAuthMiddleware())
	admin.Use(middleware.AdminMiddleware())

	{
		admin.GET("/users", controllers.GetUsers)

		admin.PUT(
			"/users/:id/role",
			controllers.UpdateUserRole,
		)

		admin.PUT(
			"/users/:id/status",
			controllers.UpdateUserStatus,
		)

		admin.DELETE(
			"/users/:id",
			controllers.DeleteUser,
		)
	}
}
