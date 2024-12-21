package routes

import (
	"message/controllers"
	"message/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	message := api.Group("/message")
	//message.Get("/", controllers.CreateMessage)
	//message.Get("/:id", controllers.IndexMessage)
	message.Post("/", controllers.CreateMessage)
	//message.Put("/:id", controllers.UpdateMessage)
	//message.Delete("/:id", controllers.DeleteMessage)

}
