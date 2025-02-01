package entities

import (
	"time"
)

type Bidang struct {
	ID         string    `gorm:"type:char(36);primaryKey"`
	NamaBidang string    `gorm:"type:varchar(255);column:nama_bidang"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
}
