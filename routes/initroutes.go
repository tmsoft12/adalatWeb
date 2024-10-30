package routes

import (
	"adalat/controllers"

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
	home := app.Group("/api")
	home.Get("/home", controllers.HomePage)
	home.Get("/banner/:id", controllers.BannerGetById)
	home.Get("/news/:id", controllers.NewsGetById)
	home.Get("/employer/:id", controllers.EmployerGetById)
	home.Get("/laws/:id", controllers.LawsGetById)
	home.Get("/media/:id", controllers.MediaGetById)

	// Bellesmeler
	news := app.Group("/api")
	news.Get("/news", controllers.NewsPage)
	news.Get("/media", controllers.MediaPage)
	news.Get("/employer", controllers.EmployerPage)
	news.Get("/laws", controllers.LawsPage)
	news.Get("/about", controllers.AboutPage)

	// Chat marşrutlary
	chat := app.Group("/api/chat")

	// Adaty HTTP marşrutlary
	chat.Get("/", controllers.Chat)           // Chat maglumatlary almak
	chat.Post("/:id", controllers.CreateChat) // Täze habar döretmek

	// WebSocket marşrutlary
	chat.Get("/ws/:id", websocket.New(controllers.ChatReal)) // WebSocket arkaly hakyky wagtly habar alyş-çalşyk

	// User ID döretmek üçin marşrut
	chat.Get("/me", controllers.Me)
}
