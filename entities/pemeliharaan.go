package entities

import "time"

type Pemeliharaan struct {
	Id                   string `gorm:"primaryKey"`
	Aset_id              string `gorm:"index"`
	Tanggal_Pemeliharaan time.Time
	Deskripsi            *string `gorm:"type:text"`
	Biaya                float64
	Created_At           time.Time
	Updated_At           time.Time
}
