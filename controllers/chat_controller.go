package controllers

import (
	"adalat/database"
	"adalat/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Me(c *fiber.Ctx) error {
	id := uuid.New().String()
	return c.Status(200).JSON(fiber.Map{"user_id": id})

}

func Chat(c *fiber.Ctx) error {
	id := c.Params("id")
	var chats []models.ChatModel
	userId := fmt.Sprintf("%d", time.Now().Unix())

	if err := database.DB.Where("user_id = ?", id).Find(&chats).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   chats,
		"userId": userId,
	})
}

func CreateChat(c *fiber.Ctx) error {
	id := c.Params("id")
	var chat models.ChatModel
	if err := c.BodyParser(&chat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Veri işlenemedi",
		})
	}
	chat.User_Id = id
	chat.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := database.DB.Create(&chat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "s",
		"data":    chat,
	})
}
