package routes

import (
	"product/controllers"
	"product/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	sale := api.Group("/sale")
	sale.Get("/", controllers.ShowSale)
	sale.Get("/:id_sale", controllers.IndexSale)
	sale.Post("/", controllers.CreateSale)
	sale.Put("/:id_sale", controllers.UpdateSale)
	sale.Delete("/:id_sale", controllers.DeleteSale)

	cat := api.Group("/cat")
	cat.Get("/", controllers.ShowCat)
	cat.Get("/:id", controllers.IndexCat)
	cat.Post("/", controllers.CreateCat)
	cat.Put("/:id", controllers.UpdateCat)
	cat.Delete("/:id", controllers.DeleteCat)

	gen := api.Group("/gen")
	gen.Get("/", controllers.ShowGen)
	gen.Get("/:id", controllers.IndexGen)
	gen.Post("/", controllers.CreateGen)
	gen.Put("/:id", controllers.UpdateGen)
	gen.Delete("/:id", controllers.DeleteGen)
}
