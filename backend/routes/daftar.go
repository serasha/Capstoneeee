package routes

import "github.com/gofiber/fiber/v2"

func SetupDaftarRoutes(app *fiber.App) {
	app.Post("/api/daftar", func(c *fiber.Ctx) error {
		// logika daftar di sini
		return c.SendString("Daftar berhasil")
	})
}
