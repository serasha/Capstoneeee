package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
)

// List log aktivitas dengan pagination
func GetLogAktifitasList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var logs []models.LogAktifitas
		var total int64
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page_size", 10)

		query := db.Model(&models.LogAktifitas{}).Preload("Admin")
		query.Count(&total)
		if err := query.Order("created_at DESC").Offset((page-1)*pageSize).Limit(pageSize).Find(&logs).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data log aktivitas"})
		}
		return c.JSON(fiber.Map{
			"data": logs,
			"total": total,
			"page": page,
			"page_size": pageSize,
		})
	}
}

// Tambah log aktivitas (bisa dipanggil dari handler lain)
func CreateLogAktifitas(db *gorm.DB, adminID uint, aksi, target, deskripsi string) error {
	log := models.LogAktifitas{
		AdminID:   adminID,
		Aksi:      aksi,
		Target:    target,
		Deskripsi: deskripsi,
	}
	return db.Create(&log).Error
} 