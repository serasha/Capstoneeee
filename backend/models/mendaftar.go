// models/mendaftar.go
package models

import (
"time"
)

type Mendaftar struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	MasyarakatID   uint      `json:"id_masyarakat"`
	PendaftaranID  uint      `json:"id_pendaftaran"`
	TanggalDaftar  time.Time `json:"tanggal_pendaftar"`
	StatusDaftar   string    `json:"status_pendaftar"`
	JenisLayanan   string    `json:"jenis_layanan"`
	CaraDaftar     string    `json:"cara_pendaftar"`

	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}