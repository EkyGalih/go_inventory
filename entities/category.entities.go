package entities

import (
	"time"
)

type Category struct {
	ID           string    `json:"ID"`
	NamaKategori string    `json:"NamaKategori"`
	Deskripsi    *string   `json:"Deskripsi"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}
