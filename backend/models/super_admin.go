package models

type SuperAdmin struct {
	ID                uint   `gorm:"primaryKey;column:id_supe_admin"`
	NamaSuperAdmin    string
	Username          string
	Password          string
	EmailSuperAdmin   string
	HakAkses          string
	StatusAktif       bool
	JumlahUserDikelola int
	LogTerakhirAkses  string
}
