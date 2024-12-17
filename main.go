package main

import (
	"crud-api/db"
	"crud-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get PORT from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	// Initialize MongoDB
	db.InitMongoDB()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.TaskRoutes(router)

	// Start server
	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)
}
