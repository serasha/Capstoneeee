package handlers

import (
	"my-app/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Create
func CreateMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Masyarakat
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		if err := db.Create(&input).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan data"})
		}
		return c.JSON(input)
	}
}

// Read All
func GetAllMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var list []models.Masyarakat
		if err := db.Find(&list).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data"})
		}
		return c.JSON(list)
	}
}

// Read by ID
func GetMasyarakatByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var item models.Masyarakat
		if err := db.First(&item, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data tidak ditemukan"})
		}
		return c.JSON(item)
	}
}

// Update
func UpdateMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var item models.Masyarakat
		if err := db.First(&item, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data tidak ditemukan"})
		}

		var updateData models.Masyarakat
		if err := c.BodyParser(&updateData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Data tidak valid"})
		}

		db.Model(&item).Updates(updateData)
		return c.JSON(item)
	}
}

// Delete
func DeleteMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&models.Masyarakat{}, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus data"})
		}
		return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
	}
}
