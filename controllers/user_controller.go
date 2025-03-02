package controllers

import (
	"context"
	"lumen/backend/config"
	"lumen/backend/prisma/db"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var body struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	err := c.BodyParser(&body) // it used to parse the request body and store the result in the body
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	newUser, err := config.DB.User.CreateOne(
		db.User.Name.Set(body.Name),
		db.User.Email.Set(body.Email),
		db.User.Password.Set(body.Password),
	).Exec(context.Background()) //context.Background() is used to create an empty context
	if err != nil {
		return c.Status(500).JSON(fiber.Map{ // 500 Internal Server Error
			"error": err.Error(),
		})
	}
	return c.Status(201).JSON(newUser) // 201 is used to create a new resource
}

func GetUsers(c *fiber.Ctx) error {
	users, err := config.DB.User.FindMany().Exec(context.Background())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := config.DB.User.FindUnique(db.User.ID.Equals(id)).Exec(context.Background())

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	var body struct{
		Name string `json:"name"`
		Email string `json:"email"`
	}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	exsistingUser, err := config.DB.User.FindUnique(db.User.ID.Equals(id)).Exec(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	
	updatedUser, err := config.DB.User.FindUnique(
		db.User.ID.Equals(exsistingUser.ID)).Update(
			db.User.Name.Set(body.Name),
			db.User.Email.Set(body.Email),
		).Exec(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(updatedUser)

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	exsistingUser, err := config.DB.User.FindUnique(db.User.ID.Equals(id)).Exec(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	_, err = config.DB.User.FindUnique(db.User.ID.Equals(exsistingUser.ID)).Delete().Exec(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(204) // 204 is used to delete a resource
}