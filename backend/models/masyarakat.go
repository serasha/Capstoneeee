package models

import (
	"time"
)

type Masyarakat struct {
	ID             uint           `gorm:"primaryKey;column:id_masyarakat"`
	NamaLengkap    string         `gorm:"size:255;not null"`
	NIK            string         `gorm:"size:16;unique;not null"`
	TanggalLahir   time.Time
	KK             string         `gorm:"size:16"`
	Email          string         `gorm:"size:255;unique"`
	Alamat         string         `gorm:"type:text"`
	NomorTelepon   string         `gorm:"size:15"`
	KtpSuami       string         `gorm:"size:255"`
	KtpIstri       string         `gorm:"size:255"`
	Ijazah         string         `gorm:"size:255"`
	PasFoto        string         `gorm:"size:255"`
	SuratNikah     string         `gorm:"size:255"`
	SuratNikahFile string         `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Pendaftarans []Pendaftaran `gorm:"foreignKey:IDMasyarakat"`
}
