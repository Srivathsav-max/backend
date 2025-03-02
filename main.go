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

func main(){
	err := godotenv.Load() // it used to load the .env file
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.InitDatabase()
	defer config.DisconnectDatabase()
	app := fiber.New()

// Set default CORS origin if not provided
corsOrigins := os.Getenv("CORS_ORIGIN")
if corsOrigins == "" {
    corsOrigins = "http://localhost:3000"
}

app.Use(cors.New(cors.Config{
    AllowOrigins: corsOrigins,
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
    AllowCredentials: true,
    ExposeHeaders: "Set-Cookie",
}))
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		err = app.Listen("")
	}else{
		err = app.Listen(":" + os.Getenv("PORT"))
	}
	
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
