package entities

import "time"

type Pemeliharaan struct {
	Id                   string `gorm:"primaryKey"`
	Aset_id              string `gorm:"index"`
	Tanggal_Pemeliharaan time.Time
	Kerusakan            *string `gorm:"type:text"`
	Perbaikan            *string `gorm:"type:text"`
	Keterangan           *string `gorm:"type:text"`
	Status               string  `gorm:"type:ENUM('Proses','Selesai');DEFAULT:'Proses'"`
	Biaya                float64
	Created_At           time.Time
	Updated_At           time.Time
}
