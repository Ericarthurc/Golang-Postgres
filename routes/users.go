package routes

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	users.Get("/")
	users.Get("/me")
	users.Get("/:id")
	users.Post("/")
	users.Post("/login")
	users.Put("/me")
	users.Delete("/me")
}
