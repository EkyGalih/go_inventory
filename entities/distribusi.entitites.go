package entities

import "time"

type Distribusi struct {
	ID               string     `json:"ID"`
	AsetID           string     `json:"AsetID"`
	BidangID         string     `json:"BidangID"`
	PegawaiID        *string    `json:"PegawaiID"`
	TanggalPerolehan time.Time  `json:"TanggalPerolehan"`
	TanggalSelesai   *time.Time `json:"TanggalSelesai,omitempty"`
	Keterangan       *string    `json:"Keterangan,omitempty"`
	CreatedAt        time.Time  `json:"CreatedAt"`
	UpdatedAt        time.Time  `json:"UpdatedAt"`

	Aset    *Aset    `json:"Aset,omitempty"`
	Bidang  *Bidang  `json:"Bidang,omitempty"`
	Pegawai *Pegawai `json:"Pegawai,omitempty"`
}
