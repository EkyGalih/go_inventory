package entities

import "time"

type Pegawai struct {
	Id          string    `json:"ID"`
	Name         string    `json:"Name"`
	IdPegawai    string    `json:"IdPegawai"`
	Foto         string    `json:"Foto"`
	JenisPegawai string    `json:"JenisPegawai"`
	Jabatan      string    `json:"Jabatan"`
	BidangID     *string   `json:"BidangID"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`

	Bidang *Bidang `json:"Bidang,omitempty"`
}
