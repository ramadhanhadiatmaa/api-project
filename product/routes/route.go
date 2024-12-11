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
}
