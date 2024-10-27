package routes

import (
	"adalat/controllers"
	"adalat/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initroutes(app *fiber.App) {
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
}
