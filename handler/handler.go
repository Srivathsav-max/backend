package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/srivathsav-max/backend/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	// Use CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	// Setup API routes
	routes.SetupRoutes(app)

	// Convert Fiber to standard HTTP handler
	adaptor.FiberApp(app).ServeHTTP(w, r)
}
