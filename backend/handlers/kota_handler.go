package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
)

// GetAllKota mengambil semua data kota
func GetAllKota(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var kotas []models.Kota
		if err := db.Where("status_aktif = ?", true).Find(&kotas).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data kota",
			})
		}
		return c.JSON(kotas)
	}
}

// CreateKota menyimpan data kota baru
func CreateKota(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Kota
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data kota",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdateKota memperbarui data kota berdasarkan ID
func UpdateKota(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var kota models.Kota

		if err := db.First(&kota, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data kota tidak ditemukan",
			})
		}

		if err := c.BodyParser(&kota); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&kota).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data kota",
			})
		}

		return c.JSON(kota)
	}
}

// DeleteKota menghapus data kota berdasarkan ID (soft delete)
func DeleteKota(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var kota models.Kota

		if err := db.First(&kota, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data kota tidak ditemukan",
			})
		}

		kota.StatusAktif = false
		if err := db.Save(&kota).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data kota",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data kota berhasil dihapus",
		})
	}
}