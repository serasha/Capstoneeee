// models/dikelola.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Dikelola struct {
	ID                  uint           `gorm:"primaryKey" json:"id_dikelola"`
	PendaftaranID       uint           `json:"id_pendaftaran"`
	StatusPengelolaan   string         `json:"status_pengelolaan"`
	CatatanPengelolaan  string         `json:"catatan_pengelolaan"`
	WaktuVerifikasi     time.Time      `json:"waktu_verifikasi"`
	Pendaftaran         Pendaftaran    `gorm:"foreignKey:PendaftaranID" json:"pendaftaran"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}