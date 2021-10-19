package routes

import (
	"eric/handlers"
	"eric/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	// users.Get("/")
	users.Get("/me", middleware.AuthenticationMiddleware, handlers.GetMe)
	// users.Get("/:id")
	users.Post("/", handlers.CreateUserHandler)
	users.Post("/login", handlers.LoginUserHandler)
	users.Post("/logout", middleware.AuthenticationMiddleware, handlers.LogoutUserHandler)
	// users.Put("/me")
	// users.Delete("/me")
}
