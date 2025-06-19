// models/admin.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Admin struct {
	ID             uint           `gorm:"primaryKey" json:"id_admin"`
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	NamaAdmin      string         `json:"nama_admin"`
	EmailAdmin     string         `json:"email_admin"`
	StatusAktif    bool           `json:"status_aktif"`
	TanggalDibuat  time.Time      `json:"tanggal_dibuat"`
	Mengelola      []Mengelola    `gorm:"foreignKey:AdminID" json:"mengelola"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}