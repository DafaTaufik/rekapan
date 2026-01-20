package main

import (
	"os"
	"rekap-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Initiate Gin
	r := gin.Default()

	// Simple route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "Server is running!",
		})
	})

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Server run
	r.Run(":" + port)
}