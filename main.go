package main

import (
	"adalat/database"
	"adalat/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDB()

	// Fiber app döretmek
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://216.150.14.12, http://127.0.0.1", // İzin verilen IP adresi
		AllowHeaders: "*",                                      // Tüm header'lara izin veriliyor
	}))

	// IP adresini almak
	ip := os.Getenv("BASE_URL")
	if ip == "" {
		ip = "localhost"
	}

	routes.Initroutes(app)

	// Serweri işledýäris
	app.Listen(":3000")
}
