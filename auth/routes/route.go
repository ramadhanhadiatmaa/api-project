package routes

import (
	"auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/user")

	user.Post("/login", controllers.Login)
	user.Post("/register", controllers.Register)
	user.Post("/upload", controllers.UploadUserImage)
	user.Put("/:username", controllers.Update)
	user.Delete("/:username", controllers.Delete)
}
