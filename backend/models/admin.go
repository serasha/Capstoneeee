package models

import "time"

type Admin struct {
	ID         uint   `gorm:"primaryKey;column:id_admin"`
	NamaAdmin  string
	Username   string
	Password   string
	Email      string
	StatusAktif bool
	TanggalDibuat time.Time
}
