package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllDikelola mengambil semua data dikelola (verifikasi pendaftaran oleh admin)
func GetAllDikelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dikelolas []models.Dikelola
		if err := db.Preload("Pendaftaran").Preload("Admin").Find(&dikelolas).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data dikelola",
			})
		}
		return c.JSON(dikelolas)
	}
}

// CreateDikelola menyimpan data dikelola baru ke database
func CreateDikelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Dikelola
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data dikelola",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdateDikelola memperbarui data dikelola berdasarkan ID
func UpdateDikelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var dikelola models.Dikelola

		if err := db.First(&dikelola, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data dikelola tidak ditemukan",
			})
		}

		if err := c.BodyParser(&dikelola); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&dikelola).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data dikelola",
			})
		}

		return c.JSON(dikelola)
	}
}

// DeleteDikelola menghapus data dikelola berdasarkan ID
func DeleteDikelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var dikelola models.Dikelola

		if err := db.First(&dikelola, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data dikelola tidak ditemukan",
			})
		}

		if err := db.Delete(&dikelola).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data dikelola",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data dikelola berhasil dihapus",
		})
	}
}
