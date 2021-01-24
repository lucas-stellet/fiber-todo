package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lucas-stellet/fiber-todo/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	app.Use(logger.New())

	setupRoutes(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}

}
