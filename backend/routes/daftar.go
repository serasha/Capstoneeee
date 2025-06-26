package routes

import (
	"github.com/gofiber/fiber/v2"
	"my-app/backend/database"
	"my-app/backend/models"
)

func SetupDaftarRoutes(route fiber.Router) {
	route.Post("/daftar", func(c *fiber.Ctx) error {
		var masyarakat models.Masyarakat

		if err := c.BodyParser(&masyarakat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Data tidak valid",
			})
		}

		if masyarakat.Email == "" || masyarakat.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email dan password wajib diisi",
			})
		}

		// Simpan ke database
		if err := database.DB.Create(&masyarakat).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan data",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Pendaftaran berhasil",
			"user":    masyarakat,
		})
	})
}
