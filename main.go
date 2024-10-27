package main

import (
	"adalat/database"
	"adalat/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(logger.New())
	ip := os.Getenv("BASE_URL")
	routes.Initroutes(app)
	app.Listen(ip + ":3000")
}
