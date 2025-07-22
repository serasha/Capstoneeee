package models

import (
	"time"
	"gorm.io/gorm"
)

type Timeline struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PendaftaranID uint           `json:"pendaftaran_id"`
	Tahap         string         `json:"tahap"` // pendaftaran, verifikasi_dokumen, seleksi_administrasi, wawancara, hasil_akhir
	Status        string         `json:"status"` // pending, in_progress, completed, rejected
	Tanggal       time.Time      `json:"tanggal"`
	Keterangan    string         `json:"keterangan"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Pendaftaran   Pendaftaran    `gorm:"foreignKey:PendaftaranID" json:"pendaftaran"`
}