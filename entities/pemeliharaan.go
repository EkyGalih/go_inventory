package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pemeliharaan struct {
	gorm.Model
	Id                   string  `gorm:"primaryKey"`
	Aset_id              string  `gorm:"index"`
	Tanggal_Pemeliharaan time.Time
	Kerusakan            *string `gorm:"type:text"`
	Perbaikan            *string `gorm:"type:text"`
	Keterangan           *string `gorm:"type:text"`
	Status               string  `gorm:"type:ENUM('Proses','Selesai');DEFAULT:'Proses'"`
	Nota                 *string `gorm:"type:text"`
	Biaya                float64
	Nama_Aset            string
	Kode_Aset            string
	Path                 *string `gorm:"type:text"`
}
