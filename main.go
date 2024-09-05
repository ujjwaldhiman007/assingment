package main

import (
	"gofiber-app/config"
	"gofiber-app/internal/app/db"
	"gofiber-app/internal/app/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Connect to MongoDB
	db.ConnectDB()

	// Set up routes
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
