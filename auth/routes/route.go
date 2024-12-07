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

	/* user.Get("/", controllers.ShowUs) */
	/* user.Get("/:username", controllers.IndexUs) */

	/* user.Put("/:username", controllers.UpdateUs)
	user.Delete("/:username", controllers.DeleteUs) */
}
