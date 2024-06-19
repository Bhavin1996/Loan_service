package main

import (
	"log"

	"loan-service/internal/mailer"
	"loan-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes
	routes.Setup(router)

	// Initialize SMTP mailer
	mailer.InitMailer("smtp.example.com", "587", "your-email@example.com", "your-email-password")

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
