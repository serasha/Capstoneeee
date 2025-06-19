// models/pendaftaran.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Pendaftaran struct {
	ID                    uint           `gorm:"primaryKey" json:"id_pendaftaran"`
	NamaPendaftar         string         `json:"nama_pendaftar"`
	AlamatPendaftar       string         `json:"alamat_pendaftar"`
	WaktuPendaftaran      time.Time      `json:"waktu_pendaftaran"`
	TanggalPendaftar      time.Time      `json:"tanggal_pendaftar"`
	StatusPendaftar       string         `json:"status_pendaftar"`
	JenisLayanan          string         `json:"jenis_layanan"`
	CaraPendaftar         string         `json:"cara_pendaftar"`
	DokumenAdministrasi   string         `json:"dokumen_administrasi_pendaftar"`
	MasyarakatID          uint           `json:"id_masyarakat"`
	Masyarakat            Masyarakat     `gorm:"foreignKey:MasyarakatID" json:"masyarakat"`
	Dikelola              []Dikelola     `gorm:"foreignKey:PendaftaranID" json:"dikelola"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
}
