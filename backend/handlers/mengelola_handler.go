package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllMengelola mengambil semua data pengelolaan (mengelola)
func GetAllMengelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mengelolas []models.Mengelola
		if err := db.Preload("Admin").Find(&mengelolas).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data pengelolaan",
			})
		}
		return c.JSON(mengelolas)
	}
}

// CreateMengelola menyimpan data pengelolaan baru
func CreateMengelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Mengelola
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data pengelolaan",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdateMengelola memperbarui data pengelolaan berdasarkan ID
func UpdateMengelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var mengelola models.Mengelola

		if err := db.First(&mengelola, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data pengelolaan tidak ditemukan",
			})
		}

		if err := c.BodyParser(&mengelola); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&mengelola).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data pengelolaan",
			})
		}

		return c.JSON(mengelola)
	}
}

// DeleteMengelola menghapus data pengelolaan berdasarkan ID
func DeleteMengelola(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var mengelola models.Mengelola

		if err := db.First(&mengelola, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data pengelolaan tidak ditemukan",
			})
		}

		if err := db.Delete(&mengelola).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data pengelolaan",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data pengelolaan berhasil dihapus",
		})
	}
}
