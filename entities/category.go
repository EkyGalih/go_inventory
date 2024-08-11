package entities

import "time"

type Category struct {
	Id            string `gorm:"primaryKey"`
	Nama_Kategori string
	Deskripsi     *string `gorm:"type:text"`
	Created_At    time.Time
	Updated_At    time.Time
}
