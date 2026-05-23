package routes

import (
    "github.com/gin-gonic/gin"

    "your-project/controllers"
    "your-project/middlewares"
)

func AdminRoutes(r *gin.RouterGroup) {

    admin := r.Group("/admin")

    admin.Use(middlewares.JWTMiddleware())
    admin.Use(middlewares.AdminMiddleware())

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