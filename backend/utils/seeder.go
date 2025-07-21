package utils

import (
	"log"
	"time"
	"gorm.io/gorm"
	"my-app/backend/models"
)

func SeedAll(db *gorm.DB) {
	// Admin
	admin := models.Admin{Username: "admin1", Password: "adminpass", EmailAdmin: "admin1@email.com", Role: "admin"}
	db.Create(&admin)

	// User
	user := models.User{Username: "user1", Password: "userpass", Role: "user"}
	db.Create(&user)

	// Masyarakat
	masyarakat := models.Masyarakat{NamaLengkap: "Budi Santoso", NIK: "1234567890123456", Email: "budi@email.com", Password: "budipass", TglPendaftaran: time.Now()}
	db.Create(&masyarakat)

	// Pendaftaran
	pendaftaran := models.Pendaftaran{
		NamaPendaftar:   "Budi Santoso",
		AlamatPendaftar: "Jl. Merdeka 1",
		WaktuPendaftaran: time.Now(),
		TanggalPendaftar: time.Now(),
		StatusPendaftar:  "pending",
		JenisLayanan:     "Transmigrasi",
		CaraPendaftar:    "online",
		DokumenAdministrasi: "-",
		MasyarakatID: masyarakat.ID,
	}
	db.Create(&pendaftaran)

	// Log Aktivitas
	logAktif := models.LogAktifitas{
		AdminID:   admin.ID,
		Aksi:      "Verifikasi",
		Target:    "Pendaftaran",
		Deskripsi: "Pendaftaran ID: 1",
		CreatedAt: time.Now(),
	}
	db.Create(&logAktif)

	log.Println("Seeder selesai dijalankan!")
} 