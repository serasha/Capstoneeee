package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
	"time"
)

// GetWawancaraByPendaftaran mengambil data wawancara berdasarkan ID pendaftaran
func GetWawancaraByPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pendaftaranID := c.Params("pendaftaran_id")
		var wawancara models.Wawancara

		if err := db.Where("pendaftaran_id = ?", pendaftaranID).Preload("Admin").First(&wawancara).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data wawancara tidak ditemukan",
			})
		}

		return c.JSON(wawancara)
	}
}

// CreateWawancara membuat jadwal wawancara baru
func CreateWawancara(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Wawancara
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		// Set admin ID dari context (dari middleware auth)
		adminID, ok := c.Locals("admin_id").(uint)
		if ok && adminID != 0 {
			input.AdminID = adminID
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal membuat jadwal wawancara",
			})
		}

		// Update timeline
		CreateTimeline(db, input.PendaftaranID, "wawancara", "in_progress", "Jadwal wawancara telah dibuat")

		return c.Status(201).JSON(input)
	}
}

// UpdateHasilWawancara memperbarui hasil wawancara
func UpdateHasilWawancara(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var wawancara models.Wawancara

		if err := db.First(&wawancara, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Data wawancara tidak ditemukan",
			})
		}

		var input struct {
			StatusWawancara  string `json:"status_wawancara"`
			HasilWawancara   string `json:"hasil_wawancara"`
			CatatanWawancara string `json:"catatan_wawancara"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		wawancara.StatusWawancara = input.StatusWawancara
		wawancara.HasilWawancara = input.HasilWawancara
		wawancara.CatatanWawancara = input.CatatanWawancara

		if err := db.Save(&wawancara).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui hasil wawancara",
			})
		}

		// Update timeline berdasarkan hasil
		var timelineStatus, keterangan string
		if input.HasilWawancara == "lulus" {
			timelineStatus = "completed"
			keterangan = "Wawancara selesai - LULUS"
		} else if input.HasilWawancara == "tidak_lulus" {
			timelineStatus = "rejected"
			keterangan = "Wawancara selesai - TIDAK LULUS"
		} else {
			timelineStatus = "completed"
			keterangan = "Wawancara selesai"
		}

		CreateTimeline(db, wawancara.PendaftaranID, "wawancara", timelineStatus, keterangan)

		return c.JSON(wawancara)
	}
}

// GetAllWawancara mengambil semua data wawancara untuk admin
func GetAllWawancara(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var wawancaras []models.Wawancara
		
		if err := db.Preload("Pendaftaran").Preload("Admin").Find(&wawancaras).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data wawancara",
			})
		}

		return c.JSON(wawancaras)
	}
}