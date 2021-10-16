package routes

import (
	"eric/controllers"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app *fiber.App) {
	items := app.Group("/api/v1/items")
	items.Get("/", controllers.GetItemsHandler)
	items.Get("/:id", controllers.GetItemHandler)
	items.Post("/", controllers.CreateItemHandler)
	items.Put("/:id", controllers.UpdateItemHandler)
	items.Delete("/:id", controllers.DeleteItemHandler)
}
