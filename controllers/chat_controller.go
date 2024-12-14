package controllers

import (
	"adalat/database"
	"adalat/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

var clients = make(map[string]*websocket.Conn)

func Me(c *fiber.Ctx) error {
	id := uuid.New().String()
	return c.Status(200).JSON(fiber.Map{"user_id": id})
}

// ChatHandler ulanyjynyň WebSocket baglanyşygyny açmak üçin ulanylýar
func ChatHandler(c *fiber.Ctx) error {
	// Ulanyjynyň ID-ni almak
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id gerekli"})
	}

	if websocket.IsWebSocketUpgrade(c) {
		// WebSocket baglanyşygyny açmak
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

// WebSocket ulanyjynyň baglanýan wagtynda we habarlary alýan wagtynda işledilýär
func WebSocket(c *websocket.Conn) {
	// Ulanyjy ID-sini alyň
	userID := c.Query("user_id")
	clients[userID] = c
	defer delete(clients, userID)

	// Ulanyjynyň öňki habarlaryny maglumatlar bazasyndan al
	var previousMessages []models.ChatModel
	if err := database.DB.Where("user_id = ?", userID).Find(&previousMessages).Error; err != nil {
		log.Println("Öňki habarlary almakda ýalňyşlyk:", err)
	} else {
		// Ulanyja öňki habarlary ugratmak
		for _, msg := range previousMessages {
			if err := c.WriteJSON(msg); err != nil {
				log.Println("Öňki habarlary ugratmakda ýalňyşlyk:", err)
				break
			}
		}
	}

	// Täze habarlary kabul edip we gaýtadan işleýäris
	for {
		var msg models.ChatModel
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Okuw ýalňyşlygy:", err)
			break
		}

		// Habar wagtyny kesgitläň
		msg.CreatedAt = time.Now().Format(time.RFC3339)

		// Maglumatlary maglumatlar bazasyna ýazga geçirmek
		if database.DB != nil {
			if err := database.DB.Create(&msg).Error; err != nil {
				log.Println("Maglumatlar bazasyna ýazmak ýalňyşlygy:", err)
				continue
			}
		} else {
			log.Println("Maglumat bazasy baglanyşygynyň geçirilmedigini görkezer")
		}

		// Ulanyjynyň ID-sine görä habar iberiň
		if client, ok := clients[msg.User_Id]; ok {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Ýazmak ýalňyşlygy:", err)
				break
			}
		}
	}
}
