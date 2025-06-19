// models/masyarakat.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Masyarakat struct {
	ID              uint           `gorm:"primaryKey" json:"id_masyarakat"`
	NamaLengkap     string         `json:"nama_lengkap"`
	NIK             string         `gorm:"unique" json:"nik"`
	KK              string         `json:"kk"`
	Email           string         `json:"email"`
	Alamat          string         `json:"alamat"`
	NomorTelepon    string         `json:"nomor_telepon"`
	KtpSuami        string         `json:"ktp_suami"`
	KtpIstri        string         `json:"ktp_istri"`
	SuratNikah      string         `json:"surat_nikah"`
	Ijazah          string         `json:"ijazah"`
	PasFoto         string         `json:"pas_foto"`
	TglPendaftaran  time.Time      `json:"tanggal_pendaftaran"`
	Pendaftarans    []Pendaftaran  `gorm:"foreignKey:MasyarakatID" json:"pendaftarans"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
