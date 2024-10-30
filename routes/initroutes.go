package routes

import (
	"adalat/controllers"
	"adalat/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func Initroutes(app *fiber.App) {
	// CORS konfigurasiýasy
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                            // Islegleriň hemmesine rugsat berýär
		AllowMethods: "GET,POST,PUT,DELETE",          // Isleg edilýän HTTP metodlaryna rugsat
		AllowHeaders: "Origin, Content-Type, Accept", // Başyklara rugsat
	}))

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
	chat.Get("/", controllers.Chat)
	chat.Post("/:id", controllers.CreateChat)
	chat.Get("/ws/:id", websocket.New(controllers.ChatReal))
	chat.Get("/me", controllers.Me)
}
