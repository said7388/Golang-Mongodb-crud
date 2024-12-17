package main

import (
	"crud-api/db"
	"crud-api/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Get PORT from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	// Initialize MongoDB
	db.InitMongoDB()

	// Initialize Gin router
	router := gin.Default()

	// Entry point
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task API"})
	})

	// Register routes
	routes.TaskRoutes(router)

	// Start server
	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)
}
