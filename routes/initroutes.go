package routes

import (
	"adalat/controllers"
	"adalat/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/websocket/v2"
)

func Initroutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// API marşrutlary üçin gurallar
	home := app.Group("/api", middleware.FakeUser)
	home.Get("/home", controllers.HomePage)
	home.Get("/banner/:id", controllers.BannerGetById)
	home.Get("/news/:id", controllers.NewsGetById)
	home.Get("/employer/:id", controllers.EmployerGetById)
	home.Get("/laws/:id", controllers.LawsGetById)
	home.Get("/media/:id", controllers.MediaGetById)

	news := app.Group("/api")
	news.Get("/news", controllers.NewsPage)
	news.Get("/media", controllers.MediaPage)
	news.Get("/employer", controllers.EmployerPage)
	news.Get("/laws", controllers.LawsPage)
	news.Get("/about", controllers.AboutPage)

	chat := app.Group("/api/chat")
	chat.Get("/", controllers.ChatHandler)
	chat.Get("/ws", websocket.New(controllers.WebSocket))
	chat.Get("/me", controllers.Me)

}
