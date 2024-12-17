package handler

import (
	"crud-api/db"
	"crud-api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler is the entry point for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize MongoDB (ensure it's initialized only once globally)
	db.InitMongoDB()

	// Initialize Gin router
	router := gin.Default()

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task API"})
	})

	// Register other routes
	routes.TaskRoutes(router)

	// Serve HTTP request through Gin
	router.ServeHTTP(w, r)
}
