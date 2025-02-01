package entities

import (
	"gorm.io/gorm"
)

type Pegawai struct {
	gorm.Model
	Id            string `gorm:"primaryKey"`
	Name          string
	Nip           string `gorm:"type:varchar(20);"`
	Foto          string `gorm:"type:varchar(255);"`
	JenisPegawai  string `gorm:"type:varchar(50);"`
	Jabatan       string `gorm:"type:varchar(50);"`
}
