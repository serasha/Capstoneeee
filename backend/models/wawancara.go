package models

import (
	"time"
	"gorm.io/gorm"
)

type Wawancara struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PendaftaranID uint           `json:"pendaftaran_id"`
	TanggalWawancara time.Time   `json:"tanggal_wawancara"`
	WaktuWawancara   string      `json:"waktu_wawancara"`
	Lokasi        string         `json:"lokasi"`
	StatusWawancara string       `json:"status_wawancara" gorm:"default:dijadwalkan"` // dijadwalkan, selesai, tidak_hadir
	HasilWawancara  string       `json:"hasil_wawancara"` // lulus, tidak_lulus
	CatatanWawancara string      `json:"catatan_wawancara"`
	AdminID       uint           `json:"admin_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Pendaftaran   Pendaftaran    `gorm:"foreignKey:PendaftaranID" json:"pendaftaran"`
	Admin         Admin          `gorm:"foreignKey:AdminID" json:"admin"`
}