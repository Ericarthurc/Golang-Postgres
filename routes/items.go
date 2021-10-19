package routes

import (
	"eric/handlers"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app *fiber.App) {
	items := app.Group("/api/v1/items")
	// items.Get("/search", handlers.GetItemsBySearchHandler)
	items.Get("/", handlers.GetItemsHandler)
	items.Get("/:id", handlers.GetItemHandler)
	items.Post("/", handlers.CreateItemHandler)
	items.Put("/:id", handlers.UpdateItemHandler)
	items.Delete("/:id", handlers.DeleteItemHandler)
}
