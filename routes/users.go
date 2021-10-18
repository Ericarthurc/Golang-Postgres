package routes

import (
	"eric/controllers"
	"eric/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	// users.Get("/")
	users.Get("/me", middleware.AuthenticationMiddleware, controllers.GetMe)
	// users.Get("/:id")
	users.Post("/", controllers.CreateUserHandler)
	users.Post("/login", controllers.LoginUserHandler)
	users.Post("/logout", middleware.AuthenticationMiddleware, controllers.LogoutUserHandler)
	// users.Put("/me")
	// users.Delete("/me")
}
