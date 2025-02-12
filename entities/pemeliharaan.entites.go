package entities

import "time"

type Pemeliharaan struct {
	ID                  string    `json:"ID"`
	AsetID              string    `json:"AsetID"`
	TanggalPemeliharaan time.Time `json:"TanggalPemeliharaan"`
	Kerusakan           *string   `json:"Kerusakan"`
	Perbaikan           *string   `json:"Perbaikan"`
	Keterangan          *string   `json:"Keterangan"`
	Status              string    `json:"Status"`
	Nota                *string   `json:"Nota"`
	Biaya               float64   `json:"Biaya"`
	CreatedAt           time.Time `json:"CreatedAt"`
	UpdatedAt           time.Time `json:"UpdatedAt"`

	Aset *Aset `json:"Aset,omitempty"`
}
