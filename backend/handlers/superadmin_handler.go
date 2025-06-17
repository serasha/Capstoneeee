package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
)

func GetAllSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data []models.SuperAdmin
		if err := db.Find(&data).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
		}
		return c.JSON(data)
	}
}

func CreateSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.SuperAdmin
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
		}
		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data"})
		}
		return c.JSON(input)
	}
}
