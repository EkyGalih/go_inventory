package entities

import "time"

type LokasiAset struct {
	Id                string `gorm:"primaryKey"`
	Aset_id           string
	Bidang_id         string
	Pegawai_id        string
	Tanggal_Perolehan time.Time
	Tanggal_Selesai   *time.Time
	Jenis_Pemanfaatan *string `gorm:"type:ENUM('Habis Pakai','Tetap');DEFAULT:'Tetap'"`
	Keterangan        *string `gorm:"type:text"`
	Created_At        time.Time
	Updated_At        time.Time
	Nama_Aset         string
	Kode_Aset         string
	Path              *string `gorm:"type:text"`
	Nama_Bidang       string
	Nama_Pegawai      string
	Nip_Pegawai       string
	Foto_Pegawai      *string `gorm:"type:text"`
	Jenis_Pegawai     string
	Jabatan           string
}
