// models/super_admin.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type SuperAdmin struct {
	ID                 uint           `gorm:"primaryKey" json:"id_super_admin"`
	NamaSuperAdmin     string         `json:"nama_super_admin"`
	Username           string         `json:"username"`
	Password           string         `json:"password"`
	EmailSuperAdmin    string         `json:"email_super_admin"`
	StatusAktif        bool           `json:"status_aktif"`
	HakAkses           string         `json:"hak_akses"`
	LogTerakhirAkses   time.Time      `json:"log_terakhir_akses"`
	JumlahUserDikelola int            `json:"jumlah_user_dikelola"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}
