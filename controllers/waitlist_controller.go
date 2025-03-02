package controllers

import (
	"context"
	"lumen/backend/config"
	"lumen/backend/prisma/db"

	"github.com/gofiber/fiber/v2"
)

func CreateWaitlist(c *fiber.Ctx) error {
	var body struct{
		Email string `json:"email"`
	}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	newWaitlist, err := config.DB.Waitlist.CreateOne(
		db.Waitlist.Email.Set(body.Email),
	).Exec(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(newWaitlist)
}

func GetWaitlist(c *fiber.Ctx) error {
	waitlist, error := config.DB.Waitlist.FindMany().Exec(context.Background())
	if error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": error.Error(),
		})
	}
	if len(waitlist) == 0{
		return c.JSON(fiber.Map{
			"message": "No Data Found",
		})
	}
	return c.JSON(waitlist)
}