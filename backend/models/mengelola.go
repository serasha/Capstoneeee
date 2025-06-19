// models/mengelola.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Mengelola struct {
	ID               uint           `gorm:"primaryKey" json:"id_mengelola"`
	AdminID          uint           `json:"id_admin"`
	CatatanAdmin     string         `json:"catatan_admin"`
	DokumenTerlampir string         `json:"dokumen_terlampir"`
	TanggalDikelola  time.Time      `json:"tanggal_dikelola"`
	Admin            Admin          `gorm:"foreignKey:AdminID" json:"admin"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}