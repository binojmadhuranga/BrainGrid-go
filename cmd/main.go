package main

import (
    "os"
    "time"

    "github.com/binojmadhuranga/BrainGrid-go/config"
    "github.com/binojmadhuranga/BrainGrid-go/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {

    godotenv.Load()

    config.ConnectDB()

    router := gin.Default()


    // CORS Configuration
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    routes.SetupRoutes(router)

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    router.Run(":" + port)
}