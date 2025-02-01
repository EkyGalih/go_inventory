package entities

import (
	"gorm.io/gorm"
)

type Tipe struct {
	Id         string  `gorm:"primaryKey"`
	Nama_Tipe  string
	Keterangan *string `gorm:"type:text"`
	gorm.Model
}
