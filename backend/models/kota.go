package models

import (
	"time"
	"gorm.io/gorm"
)

type Kota struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	NamaKota    string         `json:"nama_kota" validate:"required"`
	Provinsi    string         `json:"provinsi" validate:"required"`
	StatusAktif bool           `json:"status_aktif" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}