package models

import (
	"time"
	"gorm.io/gorm"
)

type LogAktifitas struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	AdminID    uint           `json:"admin_id"`
	Aksi       string         `json:"aksi"`
	Target     string         `json:"target"`
	Deskripsi  string         `json:"deskripsi"`
	CreatedAt  time.Time      `json:"created_at"`
	Admin      Admin          `gorm:"foreignKey:AdminID" json:"admin"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
} 