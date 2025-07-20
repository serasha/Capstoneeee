package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
	"time"
	"github.com/go-playground/validator/v10"
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

// GetPendaftaranByUser mengambil semua data pendaftaran milik user login
func GetPendaftaranByUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(float64)
		if !ok || userID == 0 {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}
		var pendaftaran []models.Pendaftaran
		if err := db.Where("masyarakat_id = ?", uint(userID)).Order("created_at DESC").Find(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pendaftaran user"})
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

		// Ambil user login dari JWT (middleware AuthRequired harus dipasang di route)
		userID, ok := c.Locals("user_id").(float64)
		if !ok || userID == 0 {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Validasi struct dengan validator
		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			fieldErrors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				jsonField := err.Field()
				switch err.Tag() {
				case "required":
					fieldErrors[jsonField] = jsonField + " wajib diisi"
				case "min":
					fieldErrors[jsonField] = jsonField + " minimal " + err.Param() + " karakter"
				case "max":
					fieldErrors[jsonField] = jsonField + " maksimal " + err.Param() + " karakter"
				case "oneof":
					fieldErrors[jsonField] = jsonField + " harus salah satu dari: " + err.Param()
				default:
					fieldErrors[jsonField] = "Format " + jsonField + " tidak valid"
				}
			}
			return c.Status(400).JSON(fiber.Map{"errors": fieldErrors})
		}

		// Set status dan waktu otomatis
		input.StatusPendaftar = "pending"
		input.WaktuPendaftaran = time.Now()
		input.TanggalPendaftar = time.Now()
		input.MasyarakatID = uint(userID)

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