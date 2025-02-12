package entities

import "time"

type Aset struct {
	ID               string    `gorm:"type:char(36);primaryKey" json:"id"`
	JenisAset        string    `gorm:"type:varchar(255)" json:"JenisAset"`
	KodeAset         string    `gorm:"type:varchar(255);unique" json:"kodeAset"`
	NamaAset         string    `gorm:"type:varchar(255)" json:"NamaAset"`
	Merek            string    `gorm:"type:varchar(255)" json:"Merek"`
	Model            string    `gorm:"type:varchar(255)" json:"Model"`
	SerialNumber     string    `gorm:"type:varchar(255)" json:"SerialNumber"`
	Deskripsi        string    `gorm:"type:text" json:"deskripsi"`
	KategoriID       string    `gorm:"type:char(36);index" json:"KategoriID"`
	TipeID           string    `gorm:"type:char(36);index" json:"TipeID"`
	TanggalPerolehan time.Time `gorm:"type:date" json:"TanggalPerolehan"`
	Status           string    `gorm:"type:varchar(255)" json:"Status"`
	Nilai            float64   `gorm:"type:decimal(15,2)" json:"Nilai"`
	Jumlah           int       `gorm:"type:int" json:"Jumlah"`
	Keterangan       string    `gorm:"type:text" json:"Keterangan"`
	Path             string    `gorm:"type:text" json:"Path"`
	Satuan           string    `gorm:"type:varchar(255)" json:"Satuan"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`

	// relasi ke category dan tipe
	Kategori *Category `json:"Kategori,omitempty"`
	Tipe     *Tipe     `json:"tipe,omitempty"`
}
