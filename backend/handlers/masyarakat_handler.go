package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
	"log"
)

func GetAllMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var masyarakat []models.Masyarakat
		if err := db.Find(&masyarakat).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
		}
		return c.JSON(masyarakat)
	}
}

func CreateMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Masyarakat
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
		}

		// Validasi dasar
		if input.NamaLengkap == "" || len(input.NIK) != 16 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi dan NIK harus 16 digit",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			log.Println("DB Error:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data"})
		}
		return c.JSON(input)
	}
}

func UpdateMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var masyarakat models.Masyarakat
		if err := db.First(&masyarakat, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
		}

		var input models.Masyarakat
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
		}

		if input.NamaLengkap == "" || len(input.NIK) != 16 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi dan NIK harus 16 digit",
			})
		}

		input.ID = masyarakat.ID
		if err := db.Save(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui data"})
		}
		return c.JSON(input)
	}
}
func DeleteMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var masyarakat models.Masyarakat
		if err := db.First(&masyarakat, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
		}
		if err := db.Delete(&masyarakat).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data"})
		}
		return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
	}
}	