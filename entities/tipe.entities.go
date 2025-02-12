package entities

import "time"

type Tipe struct {
	ID         string    `json:"ID"`
	Nama_Tipe   string    `json:"Nama_Tipe"`
	Keterangan *string   `json:"Keterangan"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}
