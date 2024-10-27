package controllers

import (
	"adalat/database"
	"adalat/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewsPage(c *fiber.Ctx) error {
	// Sahypa we limit parametrlerini al
	pageParam := c.Query("page", "1")    // Adaty sahypa: 1
	limitParam := c.Query("limit", "10") // Adaty limit: 10

	// Parametrleri integer görnüşine geçir
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1 // Adaty sahypa
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10 // Adaty limit
	}

	// Umumy sany hasapla we başlangyç nokady (offset) kesgitle
	var total int64
	offset := (page - 1) * limit

	// Täzelikleri al
	var news []models.News
	if err := database.DB.Model(&models.News{}).Count(&total).Limit(limit).Offset(offset).Find(&news).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Täzelikleri almakda säwlik"})
	}

	// BASE_URL we PORT üçin deslapky bahalary kesgitläň
	ip := os.Getenv("BASE_URL")
	if ip == "" {
		ip = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	api := "api/admin"

	for i := range news {
		news[i].Image = fmt.Sprintf("http://%s:%s/%s/%s", ip, port, api, news[i].Image)
	}

	// Sahypalaýyş maglumatlary bilen yzyna gaýtaryş
	return c.Status(200).JSON(fiber.Map{
		"data":       news,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // Umumy sahypalaryň sany
	})
}
