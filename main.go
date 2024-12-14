package main

import (
	"adalat/database"
	"adalat/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Maglumat bazasyna baglanyşmak
	database.ConnectDB()

	// Fiber app döretmek
	app := fiber.New()
	app.Use(logger.New())

	// IP adresini almak
	ip := os.Getenv("BASE_URL")
	if ip == "" {
		ip = "localhost"
	}

	// Ýollary konfigurasiýa etmek
	routes.Initroutes(app)

	// Serweri işledýäris
	app.Listen(ip + ":3000")
}
