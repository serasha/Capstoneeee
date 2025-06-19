package routes

import "github.com/gofiber/fiber/v2"

func SetupLoginRoutes(app *fiber.App) {
	app.Post("/api/login", func(c *fiber.Ctx) error {
		// logika login di sini
		return c.SendString("Login berhasil")
	})
}
