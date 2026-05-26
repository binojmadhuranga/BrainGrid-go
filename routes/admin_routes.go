package routes

import (
	"github.com/binojmadhuranga/BrainGrid-go/controllers"
	"github.com/binojmadhuranga/BrainGrid-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(api *gin.RouterGroup) {

	admin := api.Group("/admin")

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