package entities

import (
	"time"
)

type Tipe struct {
	Id         string `gorm:"primaryKey"`
	Nama_Tipe  string
	Keterangan *string `gorm:"text"`
	Created_At time.Time
	Updated_At time.Time
}
