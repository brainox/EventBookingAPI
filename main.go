package main

import (
	"event-booking-api/database"
	"event-booking-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the database
	database.InitDB()

	// Initialize Gin router
	server := gin.Default()
	// Setup routes
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
