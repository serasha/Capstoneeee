package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// API
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Halo dari Fiber!",
		})
	})

	// Serve Vue build (production mode)
	app.Static("/", "../frontend/dist")

	app.Listen(":8080")
}
