package entities

import (
	"time"

	"gorm.io/gorm"
)

type AsetTik struct {
	ID               string         `gorm:"type:char(36);primaryKey"`
	JenisAset        string         `gorm:"type:varchar(255)"`
	KodeAset         string         `gorm:"type:varchar(255);unique"`
	NamaAset         string         `gorm:"type:varchar(255)"`
	Merek            string         `gorm:"type:varchar(255)"`
	Model            string         `gorm:"type:varchar(255)"`
	SerialNumber     string         `gorm:"type:varchar(255)"`
	Deskripsi        string         `gorm:"type:text"`
	KategoriID       string         `gorm:"type:char(36);index"`
	TipeID           string         `gorm:"type:char(36);index"`
	TanggalPerolehan time.Time      `gorm:"type:date"`
	Status           string         `gorm:"type:varchar(255)"`
	Nilai            float64        `gorm:"type:decimal(15,2)"`
	Jumlah           int            `gorm:"type:int"`
	Keterangan       string         `gorm:"type:text"`
	Path             string         `gorm:"type:text"`
	Gambar           string         `gorm:"type:text"`
	Satuan           string         `gorm:"type:varchar(255)"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
