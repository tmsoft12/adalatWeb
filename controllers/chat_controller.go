package controllers

import (
	"adalat/database"
	"adalat/models"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

// Her ulanyjynyň baglanyşygyny saklamak üçin map we mutex
var clients = make(map[string]map[*websocket.Conn]bool)
var mutex = sync.Mutex{}

func Me(c *fiber.Ctx) error {
	id := uuid.New().String()
	return c.Status(200).JSON(fiber.Map{"user_id": id})
}

func Chat(c *fiber.Ctx) error {
	id := c.Params("id")
	var chats []models.ChatModel

	// Ulanyjynyň habarlaryny almak üçin maglumatlar bazasy bilen habarlaşmak
	if err := database.DB.Where("user_id = ?", id).Find(&chats).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": chats,
	})
}

func ChatReal(c *websocket.Conn) {
	id := c.Params("id")

	// Baglanan ulanyjyny goşmak
	mutex.Lock()
	if clients[id] == nil {
		clients[id] = make(map[*websocket.Conn]bool)
	}
	clients[id][c] = true
	mutex.Unlock()

	// Baglanyşyk ýapylanda ulanyjyny aýyr
	defer func() {
		mutex.Lock()
		delete(clients[id], c)
		if len(clients[id]) == 0 {
			delete(clients, id)
		}
		mutex.Unlock()
		c.Close()
	}()

	// Ulanyjynyň habarlaryny başlangyçdan almak
	var chats []models.ChatModel
	if err := database.DB.Where("user_id = ?", id).Find(&chats).Error; err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"error": "%v"}`, err.Error())))
		return
	}

	// Başlangyç maglumatlary WebSocket arkaly ibermek
	initialData := map[string]interface{}{
		"data": chats,
	}

	initialDataJson, err := json.Marshal(initialData)
	if err != nil {
		fmt.Println("Error marshalling initial data:", err)
		return
	}

	if err := c.WriteMessage(websocket.TextMessage, initialDataJson); err != nil {
		fmt.Println("Başlangyç maglumatlary iberip bolmady:", err)
		return
	}

	// Ulanyjynyň her habaryny yzyna iber
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Ulanyjy baglandy:", err)
			break
		}

		// Täze habary döredýäris
		var chat models.ChatModel
		if err := json.Unmarshal(msg, &chat); err != nil {
			fmt.Println("Habar çözmekde hata:", err)
			continue
		}
		chat.User_Id = id
		chat.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

		// Habar maglumatlaryny maglumatlar bazasyna goşmak
		if err := database.DB.Where("user_id = ?", id).Find(&chats).Error; err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"error": "%v"}`, err.Error())))
			return
		}
		// Täze habary diňe degişli ulanyja ibermek
		mutex.Lock()
		messageData := map[string]interface{}{
			"type": "new_message",
			"data": chat,
		}
		messageJson, err := json.Marshal(messageData)
		if err == nil {
			for client := range clients[id] {
				client.WriteMessage(websocket.TextMessage, messageJson)
			}
		}
		mutex.Unlock()
	}
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

	// Täze habary ähli baglanan ulanyjylara ibermek
	mutex.Lock()
	messageData := map[string]interface{}{
		"type": "new_message",
		"data": chat,
	}
	messageJson, err := json.Marshal(messageData)
	if err == nil {
		for client := range clients[id] {
			client.WriteMessage(websocket.TextMessage, messageJson)
		}
	}
	mutex.Unlock()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "s",
		"data":    chat,
	})
}
