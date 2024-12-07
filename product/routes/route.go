package routes

import (
	"product/controllers"
	"product/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	type_user := api.Group("/sale")
	type_user.Get("/", controllers.ShowSale)
	type_user.Get("/:id_sale", controllers.IndexSale)
	type_user.Post("/", controllers.CreateSale)
	type_user.Put("/:id_sale", controllers.UpdateSale)
	type_user.Delete("/:id_sale", controllers.DeleteSale)
}
