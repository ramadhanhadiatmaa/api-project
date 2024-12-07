package routes

import (
	"seller/controllers"
	"seller/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	type_user := api.Group("/type")
	type_user.Get("/", controllers.ShowTy)          
	type_user.Get("/:id_type", controllers.IndexTy)      
	type_user.Post("/", controllers.CreateTy)       
	type_user.Put("/:id_type", controllers.UpdateTy)     
	type_user.Delete("/:id_type", controllers.DeleteTy)  
}
