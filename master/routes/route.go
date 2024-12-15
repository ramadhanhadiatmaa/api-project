package routes

import (
	"master/controllers"
	"master/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	type_user := api.Group("/type")
	type_user.Get("/", controllers.ShowType)
	type_user.Get("/:id", controllers.IndexType)
	type_user.Post("/", controllers.CreateType)
	type_user.Put("/:id", controllers.UpdateType)
	type_user.Delete("/:id", controllers.DeleteType)
}
