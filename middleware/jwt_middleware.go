package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header required",
            })
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")

        if len(tokenString) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token format",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}