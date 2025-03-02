package main

import (
	"log"

	"github.com/srivathsav-max/backend/config"
	"github.com/srivathsav-max/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Always try to load .env file first
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Note: No .env file found, will use environment variables")
	}

	// Initialize database with new config management
	if err := config.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.DisconnectDatabase()

	app := fiber.New()

	// Use CORS config from environment
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.CorsOrigin,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	routes.SetupRoutes(app)

	log.Printf("🚀 Starting server on port %s", config.AppConfig.Port)
	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
