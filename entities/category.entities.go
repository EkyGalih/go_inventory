package entities

import (
	"time"
)

type Category struct {
	ID           string    `gorm:"type:char(36);primaryKey"`
	NamaKategori string    `gorm:"type:varchar(255);column:nama_kategori"`
	Deskripsi    *string   `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
}
