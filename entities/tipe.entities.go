package entities

import "time"

type Tipe struct {
	ID         string    `gorm:"type:char(36);primaryKey"`
	Nama_Tipe  string    `gorm:"type:varchar(255)"`
	Keterangan *string   `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
