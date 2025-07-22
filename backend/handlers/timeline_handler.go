package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
	"time"
)

// GetTimelineByPendaftaran mengambil timeline berdasarkan ID pendaftaran
func GetTimelineByPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pendaftaranID := c.Params("pendaftaran_id")
		var timelines []models.Timeline

		if err := db.Where("pendaftaran_id = ?", pendaftaranID).Order("tanggal ASC").Find(&timelines).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data timeline",
			})
		}

		return c.JSON(timelines)
	}
}

// CreateTimeline membuat timeline baru
func CreateTimeline(db *gorm.DB, pendaftaranID uint, tahap, status, keterangan string) error {
	timeline := models.Timeline{
		PendaftaranID: pendaftaranID,
		Tahap:         tahap,
		Status:        status,
		Tanggal:       time.Now(),
		Keterangan:    keterangan,
	}

	return db.Create(&timeline).Error
}

// UpdateTimeline memperbarui status timeline
func UpdateTimeline(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var timeline models.Timeline

		if err := db.First(&timeline, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Timeline tidak ditemukan",
			})
		}

		var input struct {
			Status     string `json:"status"`
			Keterangan string `json:"keterangan"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		timeline.Status = input.Status
		timeline.Keterangan = input.Keterangan

		if err := db.Save(&timeline).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memperbarui timeline",
			})
		}

		return c.JSON(timeline)
	}
}