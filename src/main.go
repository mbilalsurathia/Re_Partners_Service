package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pack-calculator/src/controller"
)

func main() {
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowPrivateNetwork: false,
		AllowCredentials:    true,
		ExposeHeaders:       []string{"Content-Length"},
		MaxAge:              12 * 3600, // Cache the preflight request for 12 hours
	}))

	// Define the API endpoint
	router.POST("/calculate", controller.Calculate)

	router.Run(":8080")
}
