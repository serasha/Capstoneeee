package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllSuperAdmin mengambil semua data superadmin
func GetAllSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var superadmins []models.SuperAdmin
		if err := db.Find(&superadmins).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data superadmin",
			})
		}
		return c.JSON(superadmins)
	}
}

// CreateSuperAdmin menyimpan data superadmin baru
func CreateSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.SuperAdmin
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data superadmin",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdateSuperAdmin memperbarui data superadmin berdasarkan ID
func UpdateSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var superadmin models.SuperAdmin

		if err := db.First(&superadmin, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data superadmin tidak ditemukan",
			})
		}

		if err := c.BodyParser(&superadmin); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&superadmin).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data superadmin",
			})
		}

		return c.JSON(superadmin)
	}
}

// DeleteSuperAdmin menghapus data superadmin berdasarkan ID
func DeleteSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var superadmin models.SuperAdmin

		if err := db.First(&superadmin, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data superadmin tidak ditemukan",
			})
		}

		if err := db.Delete(&superadmin).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data superadmin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data superadmin berhasil dihapus",
		})
	}
}