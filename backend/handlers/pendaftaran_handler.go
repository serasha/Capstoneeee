package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllPendaftaran mengambil semua data pendaftaran dari database
func GetAllPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pendaftaran []models.Pendaftaran
		if err := db.Preload("Masyarakat").Preload("Pengelolaan").Find(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data pendaftaran",
			})
		}
		return c.JSON(pendaftaran)
	}
}

// CreatePendaftaran menyimpan data pendaftaran baru ke database
func CreatePendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Pendaftaran
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data pendaftaran",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdatePendaftaran memperbarui data pendaftaran berdasarkan ID
func UpdatePendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var pendaftaran models.Pendaftaran

		if err := db.First(&pendaftaran, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data pendaftaran tidak ditemukan",
			})
		}

		if err := c.BodyParser(&pendaftaran); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data pendaftaran",
			})
		}

		return c.JSON(pendaftaran)
	}
}

// DeletePendaftaran menghapus data pendaftaran berdasarkan ID
func DeletePendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var pendaftaran models.Pendaftaran

		if err := db.First(&pendaftaran, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data pendaftaran tidak ditemukan",
			})
		}

		if err := db.Delete(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data pendaftaran",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data pendaftaran berhasil dihapus",
		})
	}
}