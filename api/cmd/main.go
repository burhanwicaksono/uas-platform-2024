package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "your-project-path/database"
    "your-project-path/handlers"
)

func main() {
    // Load environment variables from .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    router := gin.Default()

    // Initialize database connection
    database.InitDB()

    // Routes
    router.GET("/users", handlers.GetUsers)
    router.GET("/users/:id", handlers.GetUser)
    router.POST("/users", handlers.CreateUser)
    router.PUT("/users/:id", handlers.UpdateUser)
    router.DELETE("/users/:id", handlers.DeleteUser)

    // Run the server
    port := ":8080"
    fmt.Printf("Server is running on http://localhost%s\n", port)
    router.Run(port)
}