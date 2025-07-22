package models

import (
	"time"
	"gorm.io/gorm"
)

type Dokumen struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PendaftaranID uint           `json:"pendaftaran_id"`
	JenisDokumen  string         `json:"jenis_dokumen"` // KTP, KK, Surat Nikah, Ijazah, Pas Foto
	NamaFile      string         `json:"nama_file"`
	PathFile      string         `json:"path_file"`
	StatusVerifikasi string      `json:"status_verifikasi" gorm:"default:pending"` // pending, approved, rejected
	CatatanAdmin  string         `json:"catatan_admin"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Pendaftaran   Pendaftaran    `gorm:"foreignKey:PendaftaranID" json:"pendaftaran"`
}