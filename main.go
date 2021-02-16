package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gandhi-21/SampleFiberApi/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint"
		})
	})

	//api group
	api := app.Group("/api")

	//give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))

}
