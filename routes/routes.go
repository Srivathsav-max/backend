package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/v1/api")
	userRoutes := api.Group("/users")
	waitListRoutes := api.Group("/waitlist")

	// user Routes
	userRoutes.Post("/", controllers.CreateUser)
	userRoutes.Get("/", controllers.GetUsers)
	userRoutes.Get("/:id", controllers.GetUserByID)
	userRoutes.Put("/:id", controllers.UpdateUser)
	userRoutes.Delete("/:id", controllers.DeleteUser)

	//waitList Routes
	waitListRoutes.Post("/", controllers.CreateWaitlist)
	waitListRoutes.Get("/", controllers.GetWaitlist)
}