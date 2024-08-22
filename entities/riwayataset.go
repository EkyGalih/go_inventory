package entities

import (
	"database/sql"
	"time"
)

type Riwayat struct {
	Id           string         `json:"id"`
	Aset_id      string         `json:"aset_id"`
	Nama_Aset    string         `json:"nama_aset"`
	Kode_Aset    string         `json:"kode_aset"`
	Bidang_id    string         `json:"bidang_id"`
	Nama_Bidang  string         `json:"nama_bidang"`
	Pegawai_id   string         `json:"pegawai_id"`
	Nama_Pegawai string         `json:"nama_pegawai"`
	Nip_Pegawai  sql.NullString `json:"nip_pegawai"`
	Tanggal_Aksi time.Time      `json:"tanggal_Aksi"`
	Jenis_Aksi   string         `json:"jenis_aksi"`
	Keterangan   *string        `json:"keterangan"`
	Created_At   time.Time
	Updated_At   time.Time
}
