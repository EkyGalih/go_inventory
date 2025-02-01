package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Riwayat struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	AsetID       string         `json:"aset_id"`
	NamaAset     string         `json:"nama_aset"`
	KodeAset     string         `json:"kode_aset"`
	BidangID     string         `json:"bidang_id"`
	NamaBidang   string         `json:"nama_bidang"`
	PegawaiID    string         `json:"pegawai_id"`
	NamaPegawai  string         `json:"nama_pegawai"`
	FotoPegawai  sql.NullString `json:"foto_pegawai"`
	NipPegawai   sql.NullString `json:"nip_pegawai"`
	TanggalAksi  time.Time      `json:"tanggal_aksi"`
	JenisAksi    string         `json:"jenis_aksi"`
	Keterangan   *string        `json:"keterangan"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
