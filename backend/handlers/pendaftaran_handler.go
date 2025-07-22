package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
	"time"
	"github.com/go-playground/validator/v10"
	"fmt"
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

		// Buat timeline awal
		CreateTimeline(db, input.ID, "pendaftaran", "completed", "Pendaftaran berhasil dibuat")
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

// GetPendaftaranList untuk admin, mendukung filter dan pagination
func GetPendaftaranList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pendaftarans []models.Pendaftaran
		var total int64
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page_size", 10)
		status := c.Query("status")
		nama := c.Query("nama")
		jenis := c.Query("jenis_layanan")

		query := db.Model(&models.Pendaftaran{})
		if status != "" {
			query = query.Where("status_pendaftar = ?", status)
		}
		if nama != "" {
			query = query.Where("nama_pendaftar LIKE ?", "%"+nama+"%")
		}
		if jenis != "" {
			query = query.Where("jenis_layanan = ?", jenis)
		}

		query.Count(&total)
		if err := query.Order("created_at DESC").Offset((page-1)*pageSize).Limit(pageSize).Find(&pendaftarans).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pendaftaran"})
		}
		return c.JSON(fiber.Map{
			"data": pendaftarans,
			"total": total,
			"page": page,
			"page_size": pageSize,
		})
	}
}

// VerifikasiPendaftaran untuk update status_pendaftar
func VerifikasiPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var input struct {
			StatusPendaftar string `json:"status_pendaftar"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
		}
		var pendaftaran models.Pendaftaran
		if err := db.First(&pendaftaran, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
		}
		pendaftaran.StatusPendaftar = input.StatusPendaftar
		if err := db.Save(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal update status"})
		}

		// Tambahkan log aktivitas admin
		adminID, ok := c.Locals("admin_id").(uint)
		if ok && adminID != 0 {
			action := "Verifikasi"
			if input.StatusPendaftar == "ditolak" {
				action = "Tolak"
			} else if input.StatusPendaftar == "dikembalikan" {
				action = "Kembalikan"
			}
			deskripsi := "Pendaftaran ID: " + id
			_ = CreateLogAktifitas(db, adminID, action, "Pendaftaran", deskripsi)
		}

		// Update timeline
		var timelineStatus, keterangan string
		switch input.StatusPendaftar {
		case "verifikasi":
			timelineStatus = "completed"
			keterangan = "Dokumen telah diverifikasi dan disetujui"
		case "ditolak":
			timelineStatus = "rejected"
			keterangan = "Pendaftaran ditolak"
		case "dikembalikan":
			timelineStatus = "pending"
			keterangan = "Dokumen dikembalikan untuk diperbaiki"
		default:
			timelineStatus = "pending"
			keterangan = "Status diupdate"
		}
		
		CreateTimeline(db, uint(parseUintFromString(id)), "verifikasi_dokumen", timelineStatus, keterangan)
		return c.JSON(pendaftaran)
	}
}

func parseUintFromString(s string) uint64 {
	var result uint64
	fmt.Sscanf(s, "%d", &result)
	return result
}