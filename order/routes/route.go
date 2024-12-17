package routes

import (
	"order/controllers"
	"order/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	status := api.Group("/status")
	status.Get("/", controllers.ShowStatus)
	status.Get("/:id", controllers.IndexStatus)
	status.Post("/", controllers.CreateStatus)
	status.Put("/:id", controllers.UpdateStatus)
	status.Delete("/:id", controllers.DeleteStatus)
}
