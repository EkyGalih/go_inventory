package entities

import (
	"time"

	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	ID             string    `gorm:"type:char(36);primaryKey"`
	CompanyName    string    `gorm:"type:varchar(255)"`
	CompanyAddress string    `gorm:"type:text"`
	CompanyPhone   string    `gorm:"type:varchar(20)"`
	CompanyEmail   string    `gorm:"type:varchar(50)"`
	CompanyLogo    string    `gorm:"type:text"`
	CompanyFavicon string    `gorm:"type:text"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
