package main

import (
	"fmt"
	"os"

	"github.com/srivathsav-max/backend/config"
	"github.com/srivathsav-max/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	development := os.Getenv("DEVELOPMENT")

	if development == "true" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("⚠️ Warning: No .env file found, using environment variables")
		}
	}

	// Get port with fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.InitDatabase()
	defer config.DisconnectDatabase()

	app := fiber.New()

	// Set default CORS origin if not provided
	corsOrigins := os.Getenv("CORS_ORIGIN")
	if corsOrigins == "" {
		if development == "true" {
			corsOrigins = "http://localhost:3000"
		} else {
			corsOrigins = "https://moxium.tech"
		}
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	routes.SetupRoutes(app)

	err := app.Listen(":" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
