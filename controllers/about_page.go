package controllers

import (
	"adalat/database"
	"adalat/models"

	"github.com/gofiber/fiber/v2"
)

func AboutPage(c *fiber.Ctx) error {
	id := 1
	var about models.About
	if err := database.DB.First(&about, id).Error; err != nil {
		return c.Status(500).JSON("Err", err.Error())
	}
	return c.JSON(about)
}
