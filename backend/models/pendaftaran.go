package models

import (
	"time"
)

type Pendaftaran struct {
	ID                    uint      `gorm:"primaryKey;column:id_pendaftar"`
	IDMasyarakat          uint
	NamaPendaftar         string    `gorm:"size:255"`
	AlamatPendaftar       string    `gorm:"type:text"`
	JenisLayanan          string
	CaraPendaftar         string
	StatusPendaftar       string
	TanggalPendaftar      time.Time
	WaktuPendaftaran      time.Time
	DokumenAdministrasi   string
}
