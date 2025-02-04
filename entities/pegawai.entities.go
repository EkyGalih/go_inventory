package entities

import "time"

type Pegawai struct {
	Id           string    `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(255);"`
	IdPegawai    string    `gorm:"type:varchar(20);"`
	Foto         string    `gorm:"type:varchar(255);"`
	JenisPegawai string    `gorm:"type:varchar(50);"`
	Jabatan      string    `gorm:"type:varchar(50);"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
