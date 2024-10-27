package controllers

import (
	"adalat/database"
	"adalat/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	apiPath = "api/admin"
)

// BannerResponse is the structure for the banner response
type BannerResponse struct {
	ID    uint   `json:"id"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

func formatURL(ip, port, path, file string) string {
	return fmt.Sprintf("http://%s:%s/%s/%s", ip, port, path, file)
}

func getEnvVars() (string, string) {
	ip := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	return ip, port
}

func HomePage(c *fiber.Ctx) error {
	ip, port := getEnvVars()

	var banners []models.Banner
	if err := database.DB.Find(&banners).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Bannerleri almakda säwlik"})
	}

	var activeBanners []BannerResponse
	for _, banner := range banners {
		if banner.Is_Active {
			activeBanners = append(activeBanners, BannerResponse{
				ID:    banner.ID,
				Image: formatURL(ip, port, apiPath, banner.Image),
				Link:  banner.Link,
			})
		}
	}

	var news []models.News
	if err := database.DB.Find(&news).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Täzelikleri almakda säwlik"})
	}
	for i := range news {
		news[i].Image = formatURL(ip, port, apiPath, news[i].Image)
	}

	var media []models.Media
	if err := database.DB.Find(&media).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Media maglumatlaryny almakda säwlik"})
	}
	for i := range media {
		media[i].Cover = formatURL(ip, port, fmt.Sprintf("%s/uploads/media/cover", apiPath), media[i].Cover)
		media[i].Video = formatURL(ip, port, "video", media[i].Video)
	}

	var employers []models.Employer
	if err := database.DB.Find(&employers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Işgärleri almakda säwlik"})
	}
	for i := range employers {
		employers[i].Image = formatURL(ip, port, apiPath, employers[i].Image)
	}

	return c.Status(200).JSON(fiber.Map{
		"banners":   activeBanners,
		"news":      news,
		"media":     media,
		"employers": employers,
	})
}

func BannerGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var banner models.Banner
	if err := database.DB.First(&banner, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}
	ip := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	banner.Image = formatURL(ip, port, apiPath, banner.Image)

	return c.Status(200).JSON(banner)
}

func NewsGetById(c *fiber.Ctx) error {
	id := c.Params("id")

	var news models.News
	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}
	ip := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	news.Image = formatURL(ip, port, apiPath, news.Image)

	return c.Status(200).JSON(&news)
}

func EmployerGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var employer models.Employer
	if err := database.DB.First(&employer, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}
	ip := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	employer.Image = formatURL(ip, port, apiPath, employer.Image)

	return c.Status(200).JSON(&employer)
}

func LawsGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var laws models.Laws
	if err := database.DB.First(&laws, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}
	return c.Status(200).JSON(&laws)
}

func MediaGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var media models.Media
	if err := database.DB.First(&media, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Kanunlary almakda säwlik"})
	}
	ip := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	media.Cover = formatURL(ip, port, fmt.Sprintf("%s/uploads/media/cover", apiPath), media.Cover)
	media.Video = formatURL(ip, port, "video", media.Video)

	return c.Status(200).JSON(media)
}
