package entities

import "time"

type Riwayat struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	AsetID      string    `gorm:"type:char(36);index"`
	NamaAset    string    `gorm:"type:varchar(255)"`
	KodeAset    string    `gorm:"type:varchar(255)"`
	BidangID    string    `gorm:"type:char(36);index"`
	NamaBidang  string    `gorm:"type:varchar(255)"`
	PegawaiID   string    `gorm:"type:char(36);index"`
	NamaPegawai string    `gorm:"type:varchar(255)"`
	FotoPegawai string    `gorm:"type:text"`
	IDPegawai   string    `gorm:"type:varchar(20)"`
	TanggalAksi time.Time `gorm:"type:date"`
	JenisAksi   string    `gorm:"type:varchar(255)"`
	Keterangan  string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
