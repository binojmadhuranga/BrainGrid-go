package main

import (
    "github.com/binojmadhuranga/BrainGrid-go/config"
    "github.com/binojmadhuranga/BrainGrid-go/routes"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {

    godotenv.Load()

    config.ConnectDB()

    router := gin.Default()

    routes.SetupRoutes(router)

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    router.Run(":" + port)
}