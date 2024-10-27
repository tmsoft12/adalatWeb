package controllers

import (
	"adalat/database"
	"adalat/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func MediaPage(c *fiber.Ctx) error {
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

	// Media maglumatlaryny al
	var media []models.Media
	if err := database.DB.Model(&models.Media{}).Count(&total).Limit(limit).Offset(offset).Find(&media).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Media almakda säwlik"})
	}

	// BASE_URL we PORT gurşaw üýtgeýjilerini al
	ip := os.Getenv("BASE_URL")
	if ip == "" {
		ip = "localhost" // Deslapky IP
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Deslapky port
	}

	// Video we Cover üçin URL düzedişleri
	for i := range media {
		media[i].Video = fmt.Sprintf("http://%s:%s/video/%s", ip, port, media[i].Video)
		media[i].Cover = fmt.Sprintf("http://%s:%s/api/admin/uploads/media/cover/%s", ip, port, media[i].Cover)
	}

	// Sahypalaýyş maglumatlary bilen yzyna gaýtaryş
	return c.Status(200).JSON(fiber.Map{
		"data":       media,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // Umumy sahypalaryň sany
	})
}
