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

	loc := api.Group("/loc")
	loc.Get("/", controllers.ShowLoc)
	loc.Get("/:id", controllers.IndexLoc)
	loc.Post("/", controllers.CreateLoc)
	loc.Put("/:id", controllers.UpdateLoc)
	loc.Delete("/:id", controllers.DeleteLoc)

	add := api.Group("/add")
	add.Get("/", controllers.ShowAdd)
	add.Get("/:id", controllers.IndexAdd)
	add.Post("/", controllers.CreateAdd)
	add.Put("/:id", controllers.UpdateAdd)
	add.Delete("/:id", controllers.DeleteAdd)
}
