package entities

import (
	"database/sql"
	"time"
)

type AsetTik struct {
	Id                string `gorm:"primaryKey"`
	Jenis_Aset        string `gorm:"type:ENUM('Tetap','Habis Pakai');DEFAULT:'Tetap'"`
	Kode_Aset         string
	Nama_Aset         string
	Merek             string
	Model             string
	Serial_Number     string
	Deskripsi         *string `gorm:"type:text"`
	Kategori_id       string  `gorm:"index"`
	Tipe_id           string  `gorm:"index"`
	Tanggal_Perolehan time.Time
	Status            string `gorm:"type:ENUM('Baru', 'Baik', 'Rusak','Hilang','Perbaikan');DEFAULT:'Baru'"`
	Jumlah            float64
	Nilai             float64
	Keterangan        *string `gorm:"type:text"`
	Path              *string `gorm:"type:text"`
	Gambar            *string `gorm:"type:text"`
	Satuan            sql.NullString
	Created_At        time.Time
	Updated_At        time.Time
}
