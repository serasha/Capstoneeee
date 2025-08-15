package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
	"time"
	"fmt"
)

// GetAllPendaftaran mengambil semua data pendaftaran dari database
func GetAllPendaftaran(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pendaftaran []models.Pendaftaran
		if err := db.Preload("Masyarakat").Preload("Dikelola").Find(&pendaftaran).Error; err != nil {
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
				"error": "Permintaan tidak valid: " + err.Error(),
				"message": "Pastikan format JSON benar dan Content-Type adalah application/json",
			})
		}

		// Ambil user login dari JWT (middleware AuthRequired harus dipasang di route)
		userID, ok := c.Locals("user_id").(float64)
		if !ok || userID == 0 {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Validasi minimal agar tidak gagal karena field opsional
		missing := []string{}
		if input.NamaPendaftar == "" { missing = append(missing, "nama_pendaftar") }
		if input.NIK == "" { missing = append(missing, "nik") }
		if input.JenisLayanan == "" { missing = append(missing, "jenis_layanan") }
		if input.CaraPendaftar == "" { missing = append(missing, "cara_pendaftar") }
		if len(missing) > 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Validasi minimal gagal",
				"missing_fields": missing,
			})
		}

		// Set status dan waktu otomatis
		input.StatusPendaftar = "pending"
		input.WaktuPendaftaran = time.Now()
		input.TanggalPendaftar = time.Now()
		input.MasyarakatID = uint(userID)

		// Generate nomor pendaftaran otomatis jika kosong
		if input.NomorPendaftaran == "" {
			input.NomorPendaftaran = generateNomorPendaftaran()
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data pendaftaran: " + err.Error(),
			})
		}

		// Buat timeline awal
		CreateTimeline(db, input.ID, "pendaftaran", "completed", "Pendaftaran berhasil dibuat")
		
		return c.Status(201).JSON(fiber.Map{
			"message": "Pendaftaran berhasil dibuat",
			"data": input,
		})
	}
}

// CreatePendaftaranFlexible menyimpan data pendaftaran dengan validasi yang lebih fleksibel
func CreatePendaftaranFlexible(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			NamaPendaftar         string `json:"nama_pendaftar" validate:"required,min=3,max=100"`
			NIK                   string `json:"nik" validate:"required,min=8,max=32"`
			TanggalLahir          string `json:"tanggal_lahir" validate:"required"`
			JenisKelamin          string `json:"jenis_kelamin" validate:"required"`
			NomorKK               string `json:"nomor_kk" validate:"required"`
			PendidikanTerakhir    string `json:"pendidikan_terakhir" validate:"required"`
			Pekerjaan             string `json:"pekerjaan" validate:"required"`
			AlamatPendaftar       string `json:"alamat_pendaftar" validate:"required,min=5,max=200"`
			NomorTelepon          string `json:"nomor_telepon" validate:"required"`
			Kelurahan             string `json:"kelurahan" validate:"required"`
			JumlahAnggota         models.FlexibleInt `json:"jumlah_anggota" validate:"required"`
			KodePos               models.FlexibleInt `json:"kode_pos" validate:"required"`
			Provinsi              string `json:"provinsi" validate:"required"`
			Status                string `json:"status" validate:"required"`
			Kemantren             string `json:"kemantren" validate:"required"`
			StatusKawin           string `json:"status_kawin" validate:"required"`
			Kota                  string `json:"kota" validate:"required"`
			NamaAnggotaKeluarga   string `json:"nama_anggota_keluarga" validate:"required"`
			NamaLokasi            string `json:"nama_lokasi" validate:"required"`
			PolaUsaha             string `json:"pola_usaha" validate:"required"`
			JenisLayanan          string `json:"jenis_layanan" validate:"required,min=3,max=50"`
			CaraPendaftar         string `json:"cara_pendaftar" validate:"required"`
			DokumenAdministrasi   string `json:"dokumen_administrasi_pendaftar"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid: " + err.Error(),
				"message": "Pastikan format JSON benar dan Content-Type adalah application/json",
			})
		}

		// Ambil user login dari JWT
		userID, ok := c.Locals("user_id").(float64)
		if !ok || userID == 0 {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Validasi dengan custom validation
		if input.JenisKelamin != "laki-laki" && input.JenisKelamin != "perempuan" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Validasi gagal",
				"errors": map[string]string{
					"jenis_kelamin": "jenis_kelamin harus 'laki-laki' atau 'perempuan'",
				},
			})
		}

		if input.CaraPendaftar != "online" && input.CaraPendaftar != "offline" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Validasi gagal",
				"errors": map[string]string{
					"cara_pendaftar": "cara_pendaftar harus 'online' atau 'offline'",
				},
			})
		}

		// Buat objek Pendaftaran
		pendaftaran := models.Pendaftaran{
			NamaPendaftar:         input.NamaPendaftar,
			NIK:                   input.NIK,
			TanggalLahir:          input.TanggalLahir,
			JenisKelamin:          input.JenisKelamin,
			NomorKK:               input.NomorKK,
			PendidikanTerakhir:    input.PendidikanTerakhir,
			Pekerjaan:             input.Pekerjaan,
			AlamatPendaftar:       input.AlamatPendaftar,
			NomorTelepon:          input.NomorTelepon,
			Kelurahan:             input.Kelurahan,
			JumlahAnggota:         input.JumlahAnggota,
			KodePos:               input.KodePos,
			Provinsi:              input.Provinsi,
			Status:                input.Status,
			Kemantren:             input.Kemantren,
			StatusKawin:           input.StatusKawin,
			Kota:                  input.Kota,
			NamaAnggotaKeluarga:   input.NamaAnggotaKeluarga,
			NamaLokasi:            input.NamaLokasi,
			PolaUsaha:             input.PolaUsaha,
			NomorPendaftaran:      generateNomorPendaftaran(),
			JenisLayanan:          input.JenisLayanan,
			CaraPendaftar:         input.CaraPendaftar,
			DokumenAdministrasi:   input.DokumenAdministrasi,
			StatusPendaftar:       "pending",
			WaktuPendaftaran:      time.Now(),
			TanggalPendaftar:      time.Now(),
			MasyarakatID:          uint(userID),
		}

		if err := db.Create(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data pendaftaran: " + err.Error(),
			})
		}

		// Buat timeline awal
		CreateTimeline(db, pendaftaran.ID, "pendaftaran", "completed", "Pendaftaran berhasil dibuat")
		
		return c.Status(201).JSON(fiber.Map{
			"message": "Pendaftaran berhasil dibuat",
			"data": pendaftaran,
		})
	}
}

// CreatePendaftaranTest untuk testing tanpa authentication
func CreatePendaftaranTest(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			NamaPendaftar         string `json:"nama_pendaftar" validate:"required,min=3,max=100"`
			NIK                   string `json:"nik" validate:"required,min=8,max=32"`
			TanggalLahir          string `json:"tanggal_lahir" validate:"required"`
			JenisKelamin          string `json:"jenis_kelamin" validate:"required"`
			NomorKK               string `json:"nomor_kk" validate:"required"`
			PendidikanTerakhir    string `json:"pendidikan_terakhir" validate:"required"`
			Pekerjaan             string `json:"pekerjaan" validate:"required"`
			AlamatPendaftar       string `json:"alamat_pendaftar" validate:"required,min=5,max=200"`
			NomorTelepon          string `json:"nomor_telepon" validate:"required"`
			Kelurahan             string `json:"kelurahan" validate:"required"`
			JumlahAnggota         models.FlexibleInt `json:"jumlah_anggota" validate:"required"`
			KodePos               models.FlexibleInt `json:"kode_pos" validate:"required"`
			Provinsi              string `json:"provinsi" validate:"required"`
			Status                string `json:"status" validate:"required"`
			Kemantren             string `json:"kemantren" validate:"required"`
			StatusKawin           string `json:"status_kawin" validate:"required"`
			Kota                  string `json:"kota" validate:"required"`
			NamaAnggotaKeluarga   string `json:"nama_anggota_keluarga" validate:"required"`
			NamaLokasi            string `json:"nama_lokasi" validate:"required"`
			PolaUsaha             string `json:"pola_usaha" validate:"required"`
			JenisLayanan          string `json:"jenis_layanan" validate:"required,min=3,max=50"`
			CaraPendaftar         string `json:"cara_pendaftar" validate:"required"`
			DokumenAdministrasi   string `json:"dokumen_administrasi_pendaftar"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid: " + err.Error(),
				"message": "Pastikan format JSON benar dan Content-Type adalah application/json",
			})
		}

		// Validasi dengan custom validation
		if input.JenisKelamin != "laki-laki" && input.JenisKelamin != "perempuan" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Validasi gagal",
				"errors": map[string]string{
					"jenis_kelamin": "jenis_kelamin harus 'laki-laki' atau 'perempuan'",
				},
			})
		}

		if input.CaraPendaftar != "online" && input.CaraPendaftar != "offline" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Validasi gagal",
				"errors": map[string]string{
					"cara_pendaftar": "cara_pendaftar harus 'online' atau 'offline'",
				},
			})
		}

		// Buat objek Pendaftaran
		pendaftaran := models.Pendaftaran{
			NamaPendaftar:         input.NamaPendaftar,
			NIK:                   input.NIK,
			TanggalLahir:          input.TanggalLahir,
			JenisKelamin:          input.JenisKelamin,
			NomorKK:               input.NomorKK,
			PendidikanTerakhir:    input.PendidikanTerakhir,
			Pekerjaan:             input.Pekerjaan,
			AlamatPendaftar:       input.AlamatPendaftar,
			NomorTelepon:          input.NomorTelepon,
			Kelurahan:             input.Kelurahan,
			JumlahAnggota:         input.JumlahAnggota,
			KodePos:               input.KodePos,
			Provinsi:              input.Provinsi,
			Status:                input.Status,
			Kemantren:             input.Kemantren,
			StatusKawin:           input.StatusKawin,
			Kota:                  input.Kota,
			NamaAnggotaKeluarga:   input.NamaAnggotaKeluarga,
			NamaLokasi:            input.NamaLokasi,
			PolaUsaha:             input.PolaUsaha,
			NomorPendaftaran:      generateNomorPendaftaran(),
			JenisLayanan:          input.JenisLayanan,
			CaraPendaftar:         input.CaraPendaftar,
			DokumenAdministrasi:   input.DokumenAdministrasi,
			StatusPendaftar:       "pending",
			WaktuPendaftaran:      time.Now(),
			TanggalPendaftar:      time.Now(),
			MasyarakatID:          1, // Default untuk testing
		}

		if err := db.Create(&pendaftaran).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan data pendaftaran: " + err.Error(),
			})
		}

		// Buat timeline awal
		CreateTimeline(db, pendaftaran.ID, "pendaftaran", "completed", "Pendaftaran berhasil dibuat")
		
		return c.Status(201).JSON(fiber.Map{
			"message": "Pendaftaran berhasil dibuat",
			"data": pendaftaran,
		})
	}
}

// generateNomorPendaftaran menghasilkan nomor pendaftaran otomatis
func generateNomorPendaftaran() string {
	timestamp := time.Now().Format("20060102150405")
	return "REG-" + timestamp
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
			Catatan         string `json:"catatan"`
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
			if input.Catatan != "" {
				deskripsi += " | Catatan: " + input.Catatan
			}
			_ = CreateLogAktifitas(db, adminID, action, "Pendaftaran", deskripsi)
		}

		// Update timeline dengan catatan admin bila ada
		var timelineStatus, keterangan string
		switch input.StatusPendaftar {
		case "verifikasi":
			timelineStatus = "completed"
			if input.Catatan != "" {
				keterangan = input.Catatan
			} else {
				keterangan = "Dokumen telah diverifikasi dan disetujui"
			}
		case "ditolak":
			timelineStatus = "rejected"
			if input.Catatan != "" {
				keterangan = input.Catatan
			} else {
				keterangan = "Pendaftaran ditolak"
			}
		case "dikembalikan":
			timelineStatus = "pending"
			if input.Catatan != "" {
				keterangan = input.Catatan
			} else {
				keterangan = "Dokumen dikembalikan untuk diperbaiki"
			}
		default:
			timelineStatus = "pending"
			if input.Catatan != "" {
				keterangan = input.Catatan
			} else {
				keterangan = "Status diupdate"
			}
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