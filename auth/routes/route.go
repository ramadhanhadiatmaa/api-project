package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/api", middlewares.Auth)

	user := api.Group("/user")
	user.Get("/", controllers.ShowUs)          
	user.Get("/:id_type", controllers.IndexUs)      
	user.Post("/", controllers.CreateUs)       
	user.Put("/:id_type", controllers.UpdateUs)     
	user.Delete("/:id_type", controllers.DeleteUs)  
}
