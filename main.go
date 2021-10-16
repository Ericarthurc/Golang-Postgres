package main

import (
	"eric/database"
	"eric/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

	database.DbConnect()
	defer database.DBPool.Close()

	routes.ItemRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
