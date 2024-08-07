package entities

import (
	"time"
)

type AsetTik struct {
	Id                string `gorm:"primaryKey"`
	Kode_Aset         string
	Nama_Aset         string
	Merek             string
	Model             string
	Serial_Number     string
	Deskripsi         *string `gorm:"type:text"`
	Kategori_id       string  `gorm:"index"`
	Tanggal_Perolehan time.Time
	Status            string `gorm:"type:ENUM('Baru', 'Baik', 'Rusak','Hilang','Perbaikan');DEFAULT:'Baru'"`
	Jumlah            string
	Nilai             float64
	Keterangan        *string `gorm:"type:text"`
	Path              *string `gorm:"type:text"`
	Gambar            *string `gorm:"type:text"`
	Created_At        time.Time
	Updated_At        time.Time
}
