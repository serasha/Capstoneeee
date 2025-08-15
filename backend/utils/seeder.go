package utils

import (
	"log"
	"my-app/backend/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// HashPassword untuk encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SeedAll(db *gorm.DB) {
	log.Println("Mulai seeding data...")

	// Hash passwords
	adminPassword, _ := HashPassword("admin123")
	userPassword, _ := HashPassword("user123")
	masyarakatPassword, _ := HashPassword("budi123")

	// Super Admin
	superAdmin := models.SuperAdmin{
		Username:           "superadmin",
		Password:           adminPassword,
		NamaSuperAdmin:     "Super Administrator",
		EmailSuperAdmin:    "superadmin@transmigrasi.go.id",
		StatusAktif:        true,
		HakAkses:           "full_access",
		LogTerakhirAkses:   time.Now(),
		JumlahUserDikelola: 0,
	}
	db.FirstOrCreate(&superAdmin, models.SuperAdmin{Username: "superadmin"})

	// Admin
	admin := models.Admin{
		Username:      "admin",
		Password:      adminPassword,
		NamaAdmin:     "Administrator Dinas Transmigrasi",
		EmailAdmin:    "admin@transmigrasi.go.id",
		StatusAktif:   true,
		TanggalDibuat: time.Now(),
		Role:          "admin",
	}
	db.FirstOrCreate(&admin, models.Admin{Username: "admin"})

	// User
	user := models.User{
		Username: "user",
		Password: userPassword,
		Role:     "user",
	}
	db.FirstOrCreate(&user, models.User{Username: "user"})

	// Masyarakat (Data lengkap untuk testing)
	masyarakat := models.Masyarakat{
		NamaLengkap:    "Budi Santoso",
		NIK:            "3374051234567890",
		KK:             "3374059876543210",
		Email:          "budi.santoso@email.com",
		Password:       masyarakatPassword,
		Alamat:         "Jl. Merdeka No. 123, Semarang, Jawa Tengah",
		NomorTelepon:   "081234567890",
		KtpSuami:       "ktp_suami_budi.jpg",
		KtpIstri:       "ktp_istri_siti.jpg",
		SuratNikah:     "surat_nikah_budi_siti.pdf",
		Ijazah:         "ijazah_budi.pdf",
		PasFoto:        "pas_foto_budi.jpg",
		TglPendaftaran: time.Now(),
	}
	db.FirstOrCreate(&masyarakat, models.Masyarakat{NIK: "3374051234567890"})

	// Seed Kota (Daerah tujuan transmigrasi)
	kotas := []models.Kota{
		{NamaKota: "Palangkaraya", Provinsi: "Kalimantan Tengah", StatusAktif: true},
		{NamaKota: "Palu", Provinsi: "Sulawesi Tengah", StatusAktif: true},
		{NamaKota: "Kendari", Provinsi: "Sulawesi Tenggara", StatusAktif: true},
		{NamaKota: "Manado", Provinsi: "Sulawesi Utara", StatusAktif: true},
		{NamaKota: "Makassar", Provinsi: "Sulawesi Selatan", StatusAktif: true},
		{NamaKota: "Jayapura", Provinsi: "Papua", StatusAktif: true},
		{NamaKota: "Sorong", Provinsi: "Papua Barat", StatusAktif: true},
		{NamaKota: "Merauke", Provinsi: "Papua Selatan", StatusAktif: true},
	}
	for _, kota := range kotas {
		db.FirstOrCreate(&kota, models.Kota{NamaKota: kota.NamaKota})
	}

	// Pendaftaran (Multiple examples with different statuses)
	pendaftarans := []models.Pendaftaran{
		{
			NamaPendaftar:       "Budi Santoso",
			AlamatPendaftar:     "Jl. Merdeka No. 123, Semarang, Jawa Tengah",
			WaktuPendaftaran:    time.Now().AddDate(0, 0, -10),
			TanggalPendaftar:    time.Now().AddDate(0, 0, -10),
			StatusPendaftar:     "verifikasi",
			JenisLayanan:        "Transmigrasi Umum",
			CaraPendaftar:       "online",
			DokumenAdministrasi: "dokumen_budi.zip",
			MasyarakatID:        masyarakat.ID,
		},
		{
			NamaPendaftar:       "Siti Aminah",
			AlamatPendaftar:     "Jl. Sudirman No. 456, Yogyakarta",
			WaktuPendaftaran:    time.Now().AddDate(0, 0, -5),
			TanggalPendaftar:    time.Now().AddDate(0, 0, -5),
			StatusPendaftar:     "dikembalikan",
			JenisLayanan:        "Transmigrasi Khusus",
			CaraPendaftar:       "online",
			DokumenAdministrasi: "dokumen_siti.zip",
			MasyarakatID:        masyarakat.ID,
		},
		{
			NamaPendaftar:       "Ahmad Fauzi",
			AlamatPendaftar:     "Jl. Gatot Subroto No. 789, Bantul",
			WaktuPendaftaran:    time.Now().AddDate(0, 0, -15),
			TanggalPendaftar:    time.Now().AddDate(0, 0, -15),
			StatusPendaftar:     "Selesai",
			JenisLayanan:        "Transmigrasi Umum",
			CaraPendaftar:       "online",
			DokumenAdministrasi: "dokumen_ahmad.zip",
			MasyarakatID:        masyarakat.ID,
		},
	}

	var createdPendaftarans []models.Pendaftaran
	for _, pendaftaran := range pendaftarans {
		var existing models.Pendaftaran
		result := db.Where("nama_pendaftar = ? AND alamat_pendaftar = ?", pendaftaran.NamaPendaftar, pendaftaran.AlamatPendaftar).First(&existing)
		if result.Error != nil {
			db.Create(&pendaftaran)
			createdPendaftarans = append(createdPendaftarans, pendaftaran)
		} else {
			createdPendaftarans = append(createdPendaftarans, existing)
		}
	}

	// Timeline untuk setiap pendaftaran
	timelines := []models.Timeline{
		// Timeline untuk Budi (verifikasi)
		{
			PendaftaranID: createdPendaftarans[0].ID,
			Tahap:         "pendaftaran",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -10),
			Keterangan:    "Pendaftaran berhasil dibuat",
		},
		{
			PendaftaranID: createdPendaftarans[0].ID,
			Tahap:         "verifikasi_dokumen",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -8),
			Keterangan:    "Dokumen telah diverifikasi dan disetujui oleh admin",
		},
		// Timeline untuk Siti (dikembalikan)
		{
			PendaftaranID: createdPendaftarans[1].ID,
			Tahap:         "pendaftaran",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -5),
			Keterangan:    "Pendaftaran berhasil dibuat",
		},
		{
			PendaftaranID: createdPendaftarans[1].ID,
			Tahap:         "verifikasi_dokumen",
			Status:        "rejected",
			Tanggal:       time.Now().AddDate(0, 0, -3),
			Keterangan:    "Dokumen dikembalikan: KTP istri tidak jelas, pas foto tidak sesuai standar. Silakan upload ulang dengan kualitas yang lebih baik.",
		},
		// Timeline untuk Ahmad (selesai)
		{
			PendaftaranID: createdPendaftarans[2].ID,
			Tahap:         "pendaftaran",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -15),
			Keterangan:    "Pendaftaran berhasil dibuat",
		},
		{
			PendaftaranID: createdPendaftarans[2].ID,
			Tahap:         "verifikasi_dokumen",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -12),
			Keterangan:    "Dokumen telah diverifikasi dan disetujui",
		},
		{
			PendaftaranID: createdPendaftarans[2].ID,
			Tahap:         "wawancara",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -8),
			Keterangan:    "Wawancara selesai dilaksanakan dengan hasil LULUS",
		},
		{
			PendaftaranID: createdPendaftarans[2].ID,
			Tahap:         "keputusan",
			Status:        "completed",
			Tanggal:       time.Now().AddDate(0, 0, -3),
			Keterangan:    "Keputusan akhir: DITERIMA. Proses transmigrasi dapat dilanjutkan.",
		},
	}

	for _, timeline := range timelines {
		db.FirstOrCreate(&timeline, models.Timeline{PendaftaranID: timeline.PendaftaranID, Tahap: timeline.Tahap})
	}

	// Wawancara untuk pendaftaran yang sudah selesai
	wawancara := models.Wawancara{
		PendaftaranID:    createdPendaftarans[2].ID,
		TanggalWawancara: time.Now().AddDate(0, 0, -8),
		WaktuWawancara:   "09:00 WIB",
		Lokasi:           "Kantor Dinas Transmigrasi Yogyakarta",
		StatusWawancara:  "selesai",
		HasilWawancara:   "LULUS - Calon transmigran memenuhi syarat dan siap untuk ditempatkan",
		CatatanWawancara: "Keluarga sangat antusias dan siap untuk pindah. Memiliki pengalaman pertanian yang baik.",
	}
	db.FirstOrCreate(&wawancara, models.Wawancara{PendaftaranID: createdPendaftarans[2].ID})
	// Log Aktivitas untuk demo
	logAktivitas := []models.LogAktifitas{
		{
			AdminID:   admin.ID,
			Aksi:      "Verifikasi Dokumen",
			Target:    "Pendaftaran",
			Deskripsi: "Memverifikasi dokumen pendaftaran Budi Santoso",
			CreatedAt: time.Now().AddDate(0, 0, -8),
		},
		{
			AdminID:   admin.ID,
			Aksi:      "Pengembalian Dokumen",
			Target:    "Pendaftaran",
			Deskripsi: "Mengembalikan dokumen Siti Aminah untuk perbaikan",
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
		{
			AdminID:   admin.ID,
			Aksi:      "Konfirmasi Kelulusan",
			Target:    "Wawancara",
			Deskripsi: "Mengkonfirmasi kelulusan wawancara Ahmad Fauzi",
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
	}

	for _, log := range logAktivitas {
		db.FirstOrCreate(&log, models.LogAktifitas{AdminID: log.AdminID, Aksi: log.Aksi, Target: log.Target})
	}

	log.Println("✅ Seeder selesai dijalankan!")
	log.Println("📋 Data yang dibuat:")
	log.Println("   - 1 Super Admin (username: superadmin, password: admin123)")
	log.Println("   - 1 Admin (username: admin, password: admin123)")
	log.Println("   - 1 User (username: user, password: user123)")
	log.Println("   - 1 Masyarakat (email: budi.santoso@email.com, password: budi123)")
	log.Println("   - 8 Kota tujuan transmigrasi")
	log.Println("   - 3 Pendaftaran dengan status berbeda (verifikasi, dikembalikan, selesai)")
	log.Println("   - Timeline lengkap untuk setiap pendaftaran")
	log.Println("   - 1 Data wawancara")
	log.Println("   - 3 Log aktivitas admin")
}
