package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllAdmin mengambil semua data admin
func GetAllAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var admins []models.Admin
		if err := db.Find(&admins).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data admin",
			})
		}
		return c.JSON(admins)
	}
}

// CreateAdmin menyimpan data admin baru
func CreateAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Admin
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		input.Role = "admin" // Set role admin

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data admin",
			})
		}

		return c.Status(201).JSON(input)
	}
}

// UpdateAdmin memperbarui data admin berdasarkan ID
func UpdateAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var admin models.Admin

		if err := db.First(&admin, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data admin tidak ditemukan",
			})
		}

		if err := c.BodyParser(&admin); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		if err := db.Save(&admin).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui data admin",
			})
		}

		return c.JSON(admin)
	}
}

// DeleteAdmin menghapus data admin berdasarkan ID
func DeleteAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var admin models.Admin

		if err := db.First(&admin, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data admin tidak ditemukan",
			})
		}

		if err := db.Delete(&admin).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menghapus data admin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Data admin berhasil dihapus",
		})
	}
}

// StatistikDashboardAdmin mengembalikan statistik dashboard admin
func StatistikDashboardAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var totalUsers int64
		var totalPendaftaran int64
		var terverifikasi int64
		var menunggu int64

		db.Model(&models.User{}).Count(&totalUsers)
		db.Model(&models.Pendaftaran{}).Count(&totalPendaftaran)
		db.Model(&models.Pendaftaran{}).Where("status_pendaftar = ?", "verifikasi").Count(&terverifikasi)
		db.Model(&models.Pendaftaran{}).Where("status_pendaftar = ?", "pending").Count(&menunggu)

		return c.JSON(fiber.Map{
			"total_users": totalUsers,
			"total_pendaftaran": totalPendaftaran,
			"terverifikasi": terverifikasi,
			"menunggu": menunggu,
		})
	}
}
