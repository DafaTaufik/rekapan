package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    // Initiate Gin
    r := gin.Default()
    
    // Simple route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    // Server run on port 8080
    r.Run(":8080")
}