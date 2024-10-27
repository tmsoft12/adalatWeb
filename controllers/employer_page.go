package controllers

import (
	"adalat/database"
	"adalat/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func EmployerPage(c *fiber.Ctx) error {
	// Sahypa we limit parametrlerini al
	pageParam := c.Query("page", "1")    // Adaty sahypa: 1
	limitParam := c.Query("limit", "10") // Adaty limit: 10

	// Parametrleri integer görnüşine geçir
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Umumy sany we başlangyç nokady (offset) kesgitle
	var total int64
	offset := (page - 1) * limit

	// Isgärleri almak
	var employers []models.Employer
	if err := database.DB.Model(&models.Employer{}).Count(&total).Limit(limit).Offset(offset).Find(&employers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Işgärleri almakda säwlik"})
	}

	// BASE_URL we PORT konfigurasiýasyny al
	ip := os.Getenv("BASE_URL")
	if ip == "" {
		ip = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Suratlaryň URL-lerini düzüň
	for i := range employers {
		employers[i].Image = fmt.Sprintf("http://%s:%s/api/admin/%s", ip, port, employers[i].Image)
	}

	// Sahypalaýyş maglumatlary bilen jogap
	return c.Status(200).JSON(fiber.Map{
		"data":       employers,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // Umumy sahypalaryň sany
	})
}
