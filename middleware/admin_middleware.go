package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		role := c.GetString("role")

		if role != "ADMIN" {

			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
