package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/models"
	"path/filepath"
	"fmt"
	"os"
)

// GetDokumenByPendaftaran mengambil dokumen berdasarkan ID pendaftaran
func GetDokumenByPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pendaftaranID := c.Params("pendaftaran_id")
		var dokumens []models.Dokumen

		if err := db.Where("pendaftaran_id = ?", pendaftaranID).Find(&dokumens).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data dokumen",
			})
		}

		return c.JSON(dokumens)
	}
}

// UploadDokumen mengupload dokumen
func UploadDokumen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pendaftaranID := c.FormValue("pendaftaran_id")
		jenisDokumen := c.FormValue("jenis_dokumen")

		if pendaftaranID == "" || jenisDokumen == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Pendaftaran ID dan jenis dokumen harus diisi",
			})
		}

		// Get uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "File tidak ditemukan",
			})
		}

		// Create upload directory if not exists
		uploadDir := "./uploads/dokumen"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal membuat direktori upload",
			})
		}

		// Generate unique filename
		filename := fmt.Sprintf("%s_%s_%s", pendaftaranID, jenisDokumen, file.Filename)
		filepath := filepath.Join(uploadDir, filename)

		// Save file
		if err := c.SaveFile(file, filepath); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan file",
			})
		}

		// Save to database
		dokumen := models.Dokumen{
			PendaftaranID: uint(parseUint(pendaftaranID)),
			JenisDokumen:  jenisDokumen,
			NamaFile:      file.Filename,
			PathFile:      filepath,
		}

		if err := db.Create(&dokumen).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data dokumen",
			})
		}

		return c.Status(201).JSON(dokumen)
	}
}

// VerifikasiDokumen memverifikasi dokumen
func VerifikasiDokumen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var dokumen models.Dokumen

		if err := db.First(&dokumen, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Dokumen tidak ditemukan",
			})
		}

		var input struct {
			StatusVerifikasi string `json:"status_verifikasi"`
			CatatanAdmin     string `json:"catatan_admin"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		dokumen.StatusVerifikasi = input.StatusVerifikasi
		dokumen.CatatanAdmin = input.CatatanAdmin

		if err := db.Save(&dokumen).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal memverifikasi dokumen",
			})
		}

		return c.JSON(dokumen)
	}
}

func parseUint(s string) uint64 {
	// Simple conversion, in production use strconv.ParseUint with error handling
	var result uint64
	fmt.Sscanf(s, "%d", &result)
	return result
}