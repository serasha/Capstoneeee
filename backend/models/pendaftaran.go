// models/pendaftaran.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Pendaftaran struct {
	ID                    uint           `gorm:"primaryKey" json:"id_pendaftaran"`
	NamaPendaftar         string         `json:"nama_pendaftar" validate:"required,min=3,max=100"`
	NIK                   string         `json:"nik" validate:"required,min=8,max=32"`
	TanggalLahir          string         `json:"tanggal_lahir" validate:"required"`
	JenisKelamin          string         `json:"jenis_kelamin" validate:"required,oneof=laki-laki perempuan"`
	NomorKK               string         `json:"nomor_kk" validate:"required"`
	PendidikanTerakhir    string         `json:"pendidikan_terakhir" validate:"required"`
	Pekerjaan             string         `json:"pekerjaan" validate:"required"`
	AlamatPendaftar       string         `json:"alamat_pendaftar" validate:"required,min=5,max=200"`
	NomorTelepon          string         `json:"nomor_telepon" validate:"required"`
	Kelurahan             string         `json:"kelurahan" validate:"required"`
	JumlahAnggota         FlexibleInt    `json:"jumlah_anggota" validate:"required"`
	KodePos               FlexibleInt    `json:"kode_pos" validate:"required"`
	Provinsi              string         `json:"provinsi" validate:"required"`
	Status                string         `json:"status" validate:"required"`
	Kemantren             string         `json:"kemantren" validate:"required"`
	StatusKawin           string         `json:"status_kawin" validate:"required"`
	Kota                  string         `json:"kota" validate:"required"`
	NamaAnggotaKeluarga   string         `json:"nama_anggota_keluarga" validate:"required"`
	NamaLokasi            string         `json:"nama_lokasi" validate:"required"`
	PolaUsaha             string         `json:"pola_usaha" validate:"required"`
	NomorPendaftaran      string         `json:"nomor_pendaftaran"`
	JenisLayanan          string         `json:"jenis_layanan" validate:"required,min=3,max=50"`
	CaraPendaftar         string         `json:"cara_pendaftar" validate:"required,oneof=online offline"`
	DokumenAdministrasi   string         `json:"dokumen_administrasi_pendaftar"`
	WaktuPendaftaran      time.Time      `json:"waktu_pendaftaran"`
	TanggalPendaftar      time.Time      `json:"tanggal_pendaftar"`
	StatusPendaftar       string         `json:"status_pendaftar"`
	MasyarakatID          uint           `json:"id_masyarakat"`
	Masyarakat            Masyarakat     `gorm:"foreignKey:MasyarakatID" json:"masyarakat"`
	Dikelola              []Dikelola     `gorm:"foreignKey:PendaftaranID" json:"dikelola"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
}
