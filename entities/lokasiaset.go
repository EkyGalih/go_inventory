package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type LokasiAset struct {
	ID               string         `gorm:"primaryKey"`
	AsetID           string         `gorm:"column:aset_id"`
	BidangID         string         `gorm:"column:bidang_id"`
	PegawaiID        string         `gorm:"column:pegawai_id"`
	TanggalPerolehan time.Time      `gorm:"column:tanggal_perolehan"`
	TanggalSelesai   *time.Time     `gorm:"column:tanggal_selesai"`
	JenisPemanfaatan *string        `gorm:"type:enum('Habis Pakai','Tetap');default:'Tetap';column:jenis_pemanfaatan"`
	Keterangan       *string        `gorm:"type:text;column:keterangan"`
	NamaAset         string         `gorm:"column:nama_aset"`
	KodeAset         string         `gorm:"column:kode_aset"`
	Path             *string        `gorm:"type:text;column:path"`
	NamaBidang       string         `gorm:"column:nama_bidang"`
	NamaPegawai      string         `gorm:"column:nama_pegawai"`
	NipPegawai       string         `gorm:"column:nip_pegawai"`
	FotoPegawai      sql.NullString `gorm:"type:text;column:foto_pegawai"` // Use sql.NullString here
	JenisPegawai     string         `gorm:"column:jenis_pegawai"`
	Jabatan          string         `gorm:"column:jabatan"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
