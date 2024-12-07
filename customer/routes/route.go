package routes

import (
	"customer/controllers"
	"customer/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	type_user := api.Group("/loc")
	type_user.Get("/", controllers.ShowLoc)
	type_user.Get("/:id_loc", controllers.IndexLoc)
	type_user.Post("/", controllers.CreateLoc)
	type_user.Put("/:id_loc", controllers.UpdateLoc)
	type_user.Delete("/:id_loc", controllers.DeleteLoc)
}
