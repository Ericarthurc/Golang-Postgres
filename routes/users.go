package routes

import (
	"eric/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	// users.Get("/")
	// users.Get("/me")
	// users.Get("/:id")
	users.Post("/", controllers.CreateUserHandler)
	users.Post("/login", controllers.LoginUserHandler)
	// users.Put("/me")
	// users.Delete("/me")
}
