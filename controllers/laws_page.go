package controllers

import (
	"adalat/database"
	"adalat/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LawsPage(c *fiber.Ctx) error {
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

	// Umumy sany we başlangyç nokady (offset) kesgitle
	var total int64
	offset := (page - 1) * limit

	// Kanunlary almak
	var laws []models.Laws
	if err := database.DB.Model(&models.Laws{}).Count(&total).Limit(limit).Offset(offset).Find(&laws).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}

	// Sahypalaýyş maglumatlary bilen yzyna gaýtaryş
	return c.Status(200).JSON(fiber.Map{
		"data":       laws,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // Umumy sahypalaryň sany
	})
}
